package main

import (
	"fmt"

	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
)

func main() {
	collectionProduct := &catalog.CollectionProduct{
		BaseProductInfo: catalog.BaseProductInfo{
			SKU:  "abc",
			Name: "Product ABC",
		},
		RetailerID: 1,
	}

	db, err := drivers.GetDriver("mongodb://localhost:27017", "ecomcatalogclg")
	if err != nil {
		panic(err)
	}

	db.SavePublicProduct(collectionProduct)
	collectionProductFromDB, err := db.FindPublicProductBySKUAndRetailerID("abcx", 1)

	fmt.Println(collectionProductFromDB, err)
}
