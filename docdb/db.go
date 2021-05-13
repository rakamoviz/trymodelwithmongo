package docdb

import (
	"context"
	"fmt"
	"time"

	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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

func New(connStr, database string) (*db, error) {
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

	indxList := []string{"key", "variationFamily"}

	for _, idx := range indxList {
		ctx, cancel := context.WithTimeout(context.Background(), dbConn.timeout)
		defer cancel()
		idxn := fmt.Sprintf("%s_idx", idx)
		t := true

		_, err = dbConn.productsColl.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys: bson.M{idx: 1},
			Options: &options.IndexOptions{
				Name:       &idxn,
				Background: &t,
			},
		})
	}

	return dbConn, err
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
