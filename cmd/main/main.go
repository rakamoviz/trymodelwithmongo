package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"bitbucket.org/rappinc/gohttp"
	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
	"github.com/rakamoviz/trymodelwithmongo/typealias"
)

func main() {
	db, err := drivers.GetDriver("mongodb://localhost:27017", "ecomcatalogclg")
	if err != nil {
		panic(err)
	}

	retailerID := int64(1)

	publicProductStockAmountBr := new(big.Rat)
	publicProductStockAmountBr.SetString("12.4")

	publicProductStock := &catalog.PublicProductStock{
		StoreID:    789,
		ProductSKU: "abc",
		Stock:      10,
		UnitPrice:  typealias.BigDecimalFloat64(*publicProductStockAmountBr),
		Enabled:    true,
		Variations: []catalog.PublicProductStockVariation{},
		Discounts:  []catalog.PublicProductStockDiscount{},
	}

	collectionProductFromDB, err := db.FindPublicProductBySKUAndRetailerID(
		publicProductStock.ProductSKU,
		retailerID,
	)

	collectionProductStockVariations := []catalog.CollectionProductStockVariation{}
	for _, publicProductStockVariation := range publicProductStock.Variations {
		priceDeltaDecimal128 := publicProductStockVariation.PriceDelta.BsonDecimal128()
		collectionProductStockVariations = append(
			collectionProductStockVariations, catalog.CollectionProductStockVariation{
				VariationValue: publicProductStockVariation.VariationValue,
				Stock:          publicProductStockVariation.Stock,
				PriceDelta:     priceDeltaDecimal128,
			},
		)
	}

	unitPriceDecimal128 := publicProductStock.UnitPrice.BsonDecimal128()
	collectionProductStock := &catalog.CollectionProductStock{
		StoreID:           publicProductStock.StoreID,
		ProductSKU:        publicProductStock.ProductSKU,
		RetailerID:        retailerID,
		RetailerProductID: collectionProductFromDB.AlternativeID,
		Stock:             publicProductStock.Stock,
		UnitPrice:         unitPriceDecimal128,
		Enabled:           publicProductStock.Enabled,
		Variations:        collectionProductStockVariations,
	}

	mtCatalogProductStockVariations := []catalog.EcomMtCatalogProductStockVariation{}
	for _, publicProductStockVariation := range publicProductStock.Variations {
		mtCatalogProductStockVariations = append(
			mtCatalogProductStockVariations, catalog.EcomMtCatalogProductStockVariation{
				VariationValue: publicProductStockVariation.VariationValue,
				Stock:          publicProductStockVariation.Stock,
				PriceDelta:     publicProductStockVariation.PriceDelta,
			},
		)
	}

	mtCatalogProductStock := catalog.EcomMtCatalogProductStock{
		StoreID:           publicProductStock.StoreID,
		RetailerProductID: collectionProductFromDB.AlternativeID,
		Stock:             publicProductStock.Stock,
		UnitPrice:         publicProductStock.UnitPrice,
		Enabled:           publicProductStock.Enabled,
		Variations:        mtCatalogProductStockVariations,
	}

	fmt.Println(collectionProductFromDB, collectionProductStock, mtCatalogProductStock)

	b, err := json.Marshal(mtCatalogProductStock)

	if err != nil {
		fmt.Println("xxxxxxxx", err)
	} else {
		fmt.Println("xx", string(b))
	}

	fmt.Println("___________________________________")
	fmt.Println(collectionProductStock)
	db.SaveProductStock(collectionProductStock)
	fmt.Println("___________________________________")

	ps2, err := db.FindProductStocks(map[string]interface{}{
		"retailer_id": 1, "store_id": 789, "product_sku": "abc",
	})
	fmt.Println("eeee", *ps2[0])

	restClient := gohttp.NewClient(gohttp.Options{
		ClientName: "ecom-catalog-clg",
		Timeout:    30 * time.Second,
	})

	apiClient := ecomMtCatalogAPI.New(restClient, ecomMtCatalogAPI.Option{
		Host:                "",
		Country:             "dev",
		EcomMtCatalogAPIKey: "",
		ClientName:          "ecom-catalog-clg",
	})

	err = apiClient.SyncProductStock(mtCatalogProductStock)
	if err != nil {
		fmt.Println(err)
	}
}
