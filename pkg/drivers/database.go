package drivers

import (
	"github.com/rakamoviz/trymodelwithmongo/docdb"
	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
)

type DB interface {
	SaveProduct(*catalog.ProductStore) error
	FindProduct(int64, int64) (*catalog.ProductStore, error)

	FindPublicProductBySKUAndRetailerID(string, int64) (*catalog.CollectionProduct, error)
	SavePublicProduct(*catalog.CollectionProduct) (*int64, error)

	SaveProductStock(*catalog.CollectionProductStock) error
	FindProductStock(int64, int64, string) (*catalog.CollectionProductStock, error)
	FindProductStocks(map[string]interface{}) ([]*catalog.CollectionProductStock, error)
}

func GetDriver(connStr, database string) (DB, error) {
	return docdb.New(connStr, database)
}
