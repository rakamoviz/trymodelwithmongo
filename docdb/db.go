package docdb

import (
	"context"
	"fmt"
	"time"

	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Sequence struct {
	ID            string `bson:"_id"`
	SequenceValue int64  `bson:"sequence_value"`
}

type db struct {
	conn               *mongo.Database
	productsColl       *mongo.Collection
	publicProductsColl *mongo.Collection
	migrationsColl     *mongo.Collection
	sequenceColl       *mongo.Collection
	productStocksColl  *mongo.Collection

	timeout time.Duration
}

func (d *db) incrementSequence(sequenceName string) (int64, error) {
	seq := Sequence{}

	after := options.After
	upsert := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	err := d.sequenceColl.FindOneAndUpdate(ctx, bson.M{"_id": sequenceName}, bson.M{"$inc": bson.M{"sequence_value": 1}}, &opt).Decode(&seq)

	if err != nil {
		return -1, err
	}

	return seq.SequenceValue, nil
}

const StatusPublished = "published"

type ProductStoreImage struct {
	Position int    `json:"position"`
	Path     string `json:"path"`
}

func (d *db) SavePublicProduct(p *catalog.CollectionProduct) (*int64, error) {
	logrus.Debugf("[docdb] saving: %s", p.SKU)

	key, err := d.incrementSequence("public_product")
	if err != nil {
		return nil, err
	}
	p.ID = key

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	_, err = d.publicProductsColl.InsertOne(ctx, p)
	if err != nil {
		logrus.Errorf("[docdb] SavePublicProduct: %s", err.Error())
	}

	return &key, err
}

func (d *db) SaveProduct(ps *catalog.ProductStore) error {
	if ps == nil {
		return nil
	}

	logrus.Debugf("[docdb] saving: %s [%s]", ps.Key, ps.VariationFamily)

	key := buildKey(ps.VirtualProductID, ps.VirtualStoreID)

	ps.CreatedAt = time.Now()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"key": key}
	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	_, err := d.productsColl.UpdateOne(ctx, filter, bson.M{"$set": ps}, opts)
	if err != nil {
		logrus.Errorf("[docdb] findProduct: %s\n", err.Error())
	}

	return err
}

func (d *db) SaveProductStock(
	ps catalog.ProductStockEntity,
) error {
	logrus.Debugf(
		"[docdb] saving CollectionProductStock (retailer_id: %d, store_id: %d, product_sku: %s)",
		ps.RetailerID, ps.StoreID, ps.ProductSKU,
	)

	now := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	psCopy := ps
	psCopy.CreatedAt = nil
	psCopy.UpdatedAt = &now

	updateOptions := options.Update().SetUpsert(true)
	_, err := d.productStocksColl.UpdateOne(ctx, bson.M{
		"retailer_id": psCopy.RetailerID,
		"store_id":    psCopy.StoreID,
		"product_sku": psCopy.ProductSKU,
	}, bson.M{
		"$set":         psCopy,
		"$setOnInsert": bson.M{"created_at": now, "_id": primitive.NewObjectID()},
	}, updateOptions)

	if err != nil {
		logrus.Errorf("[docdb] SaveProductStock: %s\n", err.Error())
		return &drivers.DBError{
			StatusCode: -1,
			Err:        err,
		}
	}

	return nil
}

func (d *db) FindProductStock(key catalog.ProductStockEntityKey) (catalog.ProductStockEntity, error) {
	ps := catalog.ProductStockEntity{}

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	rawFilter := map[string]interface{}{
		"retailer_id": key.RetailerID,
		"store_id":    key.StoreID,
		"product_sku": key.ProductSKU,
	}

	err := d.productStocksColl.FindOne(ctx, bson.M(rawFilter)).Decode(&ps)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logrus.Debugf("[docdb] FindProductStock miss: %s\n", rawFilter)
			return ps, &drivers.DBError{
				StatusCode: 0,
				Err:        err,
			}
		} else {
			logrus.Errorf("[docdb] FindProductStock: %s\n", err.Error())
			return ps, &drivers.DBError{
				StatusCode: -1,
				Err:        err,
			}
		}
	}

	return ps, err
}

func (d *db) FindProductStocks(rawFilter map[string]interface{}) ([]catalog.ProductStockEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	cur, err := d.productStocksColl.Find(ctx, bson.M(rawFilter))
	results := []catalog.ProductStockEntity{}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logrus.Debugf("[docdb] FindProductStock miss: %s\n", rawFilter)
			return results, &drivers.DBError{
				StatusCode: 0,
				Err:        err,
			}
		} else {
			logrus.Errorf("[docdb] FindProductStock: %s\n", err.Error())
			return results, &drivers.DBError{
				StatusCode: -1,
				Err:        err,
			}
		}
	}

	for cur.Next(context.Background()) {
		elem := catalog.ProductStockEntity{}
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Parse error : ", err)
			return nil, err
		}
		fmt.Println("Log : ", elem)
		results = append(results, elem)
	}

	return results, err
}

func (d *db) findProduct(key string) (*catalog.ProductStore, error) {
	product := catalog.ProductStore{}
	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	filter := bson.M{"key": key}
	err := d.productsColl.FindOne(ctx, filter).
		Decode(&product)

	if err == mongo.ErrNoDocuments {
		logrus.Debugf("[docdb] findProduct miss: %s\n", key)
		return nil, nil
	}

	if err != nil {
		logrus.Errorf("[docdb] findProduct: %s\n", err.Error())
	}

	return &product, err
}

func (d *db) FindProduct(storeID, productID int64) (*catalog.ProductStore, error) {
	ps, err := d.findProduct(buildKey(productID, storeID))
	if err != nil {
		logrus.Errorf("[docdb] No: %s\n", err.Error())
	}

	return ps, err
}

func (d *db) FindPublicProductBySKUAndRetailerID(sku string, retailerID int64) (*catalog.CollectionProduct, error) {
	product := catalog.CollectionProduct{}
	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	filter := bson.M{"sku": sku, "retailer_id": retailerID}
	err := d.publicProductsColl.FindOne(ctx, filter).
		Decode(&product)

	if err == mongo.ErrNoDocuments {
		logrus.Debugf("[docdb] FindPublicProductBySKUAndRetailerID miss: sku: %s, retailerID: %d", sku, retailerID)
		return nil, nil
	}

	if err != nil {
		logrus.Errorf("[docdb] FindPublicProductBySKUAndRetailerID: %s\n", err.Error())
	}

	return &product, err
}

func createProductsCollIndex(dbConn *db) error {
	indxList := []string{"key", "variationFamily"}

	for _, idx := range indxList {
		ctx, cancel := context.WithTimeout(context.Background(), dbConn.timeout)
		defer cancel()
		idxn := fmt.Sprintf("%s_idx", idx)
		t := true

		_, err := dbConn.productsColl.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys: bson.M{idx: 1},
			Options: &options.IndexOptions{
				Name:       &idxn,
				Background: &t,
			},
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func createProductStocksCollIndex(dbConn *db) error {
	indxList := []string{"retailer_id", "store_id", "product_sku"}

	for _, idx := range indxList {
		ctx, cancel := context.WithTimeout(context.Background(), dbConn.timeout)
		defer cancel()
		idxn := fmt.Sprintf("%s_idx", idx)
		t := true

		_, err := dbConn.productStocksColl.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys: bson.M{idx: 1},
			Options: &options.IndexOptions{
				Name:       &idxn,
				Background: &t,
			},
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func NewDocDb(connStr, database string) (*db, error) {
	conn, err := connection(connStr, database)
	if err != nil {
		return nil, err
	}

	to := 5 * time.Second
	dbConn, err := &db{
		conn:               conn,
		productsColl:       conn.Collection("products"),
		publicProductsColl: conn.Collection("public_products"),
		migrationsColl:     conn.Collection("catalogmigrations"),
		sequenceColl:       conn.Collection("sequence"),
		productStocksColl:  conn.Collection("product_stocks"),
		timeout:            to,
	}, nil
	if err != nil {
		return nil, err
	}

	err = createProductsCollIndex(dbConn)
	if err != nil {
		return nil, err
	}

	err = createProductStocksCollIndex(dbConn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func buildKey(pid, sid int64) string {
	return fmt.Sprintf("%d_%d", pid, sid)
}

func connection(connStr, database string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client.Database(database), nil
}

func GetDriver(connStr, database string) (drivers.DB, error) {
	return NewDocDb(connStr, database)
}
