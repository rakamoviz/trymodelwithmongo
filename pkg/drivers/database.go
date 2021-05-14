package drivers

import (
	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
)

type DBError struct {
	StatusCode int

	Err error
}

func (r *DBError) Error() string {
	return r.Err.Error()
}

type DB interface {
	SaveProduct(*catalog.ProductStore) error
	FindProduct(int64, int64) (*catalog.ProductStore, error)

	FindPublicProductBySKUAndRetailerID(string, int64) (*catalog.CollectionProduct, error)
	SavePublicProduct(*catalog.CollectionProduct) (*int64, error)

	SaveProductStock(catalog.ProductStockEntity) error
	FindProductStock(catalog.ProductStockEntityKey) (catalog.ProductStockEntity, error)
	FindProductStocks(map[string]interface{}) ([]catalog.ProductStockEntity, error)
}
