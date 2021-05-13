package main

import (
	"fmt"
	"math/big"

	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
)

func main() {
	db, err := drivers.GetDriver("mongodb://localhost:27017", "ecomcatalogclg")
	if err != nil {
		panic(err)
	}

	publicProductStockAmountBr := new(big.Rat)
	publicProductStockAmountBr.SetString("12.4")

	publicProductStock := &catalog.PublicProductStock{
		StoreID:           789,
		ProductSKU:        "abc",
		RetailerID:        1,
		RetailerProductID: 456,
		Stock:             10,
		UnitPrice:         catalog.BigDecimal(*publicProductStockAmountBr),
		Enabled:           true,
		Variations:        &[]catalog.PublicProductStockVariation{},
		Discounts:         &[]catalog.PublicProductStockDiscount{},
	}

	/*
		collectionProduct := &catalog.CollectionProduct{
			BaseProductInfo: catalog.BaseProductInfo{
				SKU:  "abc",
				Name: "Product ABC",
			},
			RetailerID: 1,
		}

		db.SavePublicProduct(collectionProduct)
		collectionProductFromDB, err := db.FindPublicProductBySKUAndRetailerID("abcx", 1)
	*/

	collectionProductFromDB, err := db.FindPublicProductBySKUAndRetailerID(
		publicProductStock.ProductSKU,
		publicProductStock.RetailerID,
	)

	fmt.Println(collectionProductFromDB)
}
