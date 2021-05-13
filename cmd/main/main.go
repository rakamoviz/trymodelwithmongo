package main

import (
	"encoding/json"
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

	retailerID := int64(1)

	publicProductStockAmountBr := new(big.Rat)
	publicProductStockAmountBr.SetString("12.4")

	publicProductStock := &catalog.PublicProductStock{
		StoreID:    789,
		ProductSKU: "abc",
		Stock:      10,
		UnitPrice:  catalog.BigDecimal(*publicProductStockAmountBr),
		Enabled:    true,
		Variations: []catalog.PublicProductStockVariation{},
		Discounts:  []catalog.PublicProductStockDiscount{},
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
		retailerID,
	)

	/*
		type CollectionProductStockVariation struct {
			VariationValue string     `json:"variation_value"`
			Stock          int32      `json:"stock"`
			PriceDelta     BigDecimal `json:"price_delta"`
		}

		type CollectionProductStockDiscount struct {
			DiscountType string     `json:"discount_type"`
			Value        string     `json:"value"`
			BeginDate    CustomTime `json:"begin_date"`
			EndDate      CustomTime `json:"end_date"`
		}
	*/
	collectionProductStockVariations := []catalog.CollectionProductStockVariation{}
	for _, publicProductStockVariation := range publicProductStock.Variations {
		collectionProductStockVariations = append(
			collectionProductStockVariations, catalog.CollectionProductStockVariation{
				VariationValue: publicProductStockVariation.VariationValue,
				Stock:          publicProductStockVariation.Stock,
				PriceDelta:     publicProductStockVariation.PriceDelta,
			},
		)
	}

	collectionProductStock := &catalog.CollectionProductStock{
		StoreID:           publicProductStock.StoreID,
		ProductSKU:        publicProductStock.ProductSKU,
		RetailerID:        retailerID,
		RetailerProductID: collectionProductFromDB.AlternativeID,
		Stock:             publicProductStock.Stock,
		UnitPrice:         publicProductStock.UnitPrice,
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

	mtCatalogProductStock := &catalog.EcomMtCatalogProductStock{
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

	db.SaveProductStock(collectionProductStock)
	ps2, err := db.FindProductStocks(map[string]interface{}{
		"retailer_id": 1, "store_id": 789, "product_sku": "abc",
	})
	fmt.Println("eeee", *ps2[0])
}
