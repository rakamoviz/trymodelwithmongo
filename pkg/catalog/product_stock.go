package catalog

import (
	"time"

	"github.com/rakamoviz/trymodelwithmongo/typealias"
)

type PublicProductStockVariation struct {
	VariationValue string                   `json:"variation_value"`
	Stock          int32                    `json:"stock"`
	PriceDelta     typealias.DecimalFloat64 `json:"price_delta"`
}

type PublicProductStockDiscount struct {
	DiscountType string         `json:"discount_type"`
	Value        string         `json:"value"`
	BeginDate    typealias.Date `json:"begin_date"`
	EndDate      typealias.Date `json:"end_date"`
}

type PublicProductStock struct {
	StoreID    int64                         `json:"store_id"`
	ProductSKU string                        `json:"product_sku"`
	Stock      int64                         `json:"stock"`
	UnitPrice  typealias.DecimalFloat64      `json:"unit_price"`
	Enabled    bool                          `json:"enabled"`
	Variations []PublicProductStockVariation `json:"variations"`
	Discounts  []PublicProductStockDiscount  `json:"discounts"`
}

type CollectionProductStockVariation struct {
	VariationValue string  `bson:"variation_value"`
	Stock          int32   `bson:"stock"`
	PriceDelta     float64 `bson:"price_delta"`
}

type CollectionProductStockDiscount struct {
	DiscountType string         `bson:"discount_type"`
	Value        string         `bson:"value"`
	BeginDate    typealias.Date `bson:"begin_date"`
	EndDate      typealias.Date `bson:"end_date"`
}

type CollectionProductStock struct {
	StoreID           int64                             `bson:"store_id"`
	ProductSKU        string                            `bson:"product_sku"`
	RetailerID        int64                             `bson:"retailer_id"`
	RetailerProductID *int64                            `bson:"retailer_product_id"`
	Stock             int64                             `bson:"stock"`
	UnitPrice         float64                           `bson:"unit_price"`
	Enabled           bool                              `bson:"enabled"`
	Variations        []CollectionProductStockVariation `bson:"variations"`
	Discounts         []CollectionProductStockDiscount  `bson:"discounts"`
	CreatedAt         *time.Time                        `bson:"created_at,omitempty"`
	UpdatedAt         *time.Time                        `bson:"updated_at,omitempty"`
}

type EcomMtCatalogProductStockVariation struct {
	RetailerProductID int64                    `json:"retailerProductId"`
	VariationValue    string                   `json:"variationValue"`
	Stock             int32                    `json:"stock"`
	PriceDelta        typealias.DecimalFloat64 `json:"priceDelta"`
}

type EcomMtCatalogProductStock struct {
	StoreID           int64                                `json:"storeId"`
	RetailerProductID *int64                               `json:"retailerProductId"`
	Stock             int64                                `json:"stock"`
	UnitPrice         typealias.DecimalFloat64             `json:"unitPrice"`
	Enabled           bool                                 `json:"enabled"`
	Variations        []EcomMtCatalogProductStockVariation `json:"variations"`
}
