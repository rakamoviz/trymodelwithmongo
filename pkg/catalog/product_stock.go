package catalog

import (
	"time"

	"github.com/rakamoviz/trymodelwithmongo/typealias"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductStockVariationDTO struct {
	VariationValue string                   `json:"variation_value"`
	Stock          int32                    `json:"stock"`
	PriceDelta     typealias.DecimalFloat64 `json:"price_delta"`
}

type ProductStockDiscountDTO struct {
	DiscountType string         `json:"discount_type"`
	Value        string         `json:"value"`
	BeginDate    typealias.Date `json:"begin_date"`
	EndDate      typealias.Date `json:"end_date"`
}

type ProductStockDTO struct {
	StoreID    int64                      `json:"store_id"`
	ProductSKU string                     `json:"product_sku"`
	Stock      int64                      `json:"stock"`
	UnitPrice  typealias.DecimalFloat64   `json:"unit_price"`
	Enabled    bool                       `json:"enabled"`
	Variations []ProductStockVariationDTO `json:"variations"`
	Discounts  []ProductStockDiscountDTO  `json:"discounts"`
}

type ProductStockVariationEntity struct {
	VariationValue string               `bson:"variation_value"`
	Stock          int32                `bson:"stock"`
	PriceDelta     primitive.Decimal128 `bson:"price_delta"`
}

type ProductStockDiscountEntity struct {
	DiscountType string         `bson:"discount_type"`
	Value        string         `bson:"value"`
	BeginDate    typealias.Date `bson:"begin_date"`
	EndDate      typealias.Date `bson:"end_date"`
}

type ProductStockEntityKey struct {
	StoreID    int64  `bson:"store_id"`
	ProductSKU string `bson:"product_sku"`
	RetailerID int64  `bson:"retailer_id"`
}

type ProductStockEntity struct {
	ID                    *primitive.ObjectID `bson:"_id,omitempty"`
	ProductStockEntityKey `bson:",inline"`
	RetailerProductID     *int64                        `bson:"retailer_product_id"`
	Stock                 int64                         `bson:"stock"`
	UnitPrice             primitive.Decimal128          `bson:"unit_price"`
	Enabled               bool                          `bson:"enabled"`
	Variations            []ProductStockVariationEntity `bson:"variations"`
	Discounts             []ProductStockDiscountEntity  `bson:"discounts"`
	CreatedAt             *time.Time                    `bson:"created_at,omitempty"`
	UpdatedAt             *time.Time                    `bson:"updated_at,omitempty"`
}

type EcomMtCatalogProductStockVariationDTO struct {
	RetailerProductID int64                    `json:"retailerProductId"`
	VariationValue    string                   `json:"variationValue"`
	Stock             int32                    `json:"stock"`
	PriceDelta        typealias.DecimalFloat64 `json:"priceDelta"`
}

type EcomMtCatalogProductStockDTO struct {
	StoreID           int64                                   `json:"storeId"`
	RetailerProductID *int64                                  `json:"retailerProductId"`
	Stock             int64                                   `json:"stock"`
	UnitPrice         typealias.DecimalFloat64                `json:"unitPrice"`
	Enabled           bool                                    `json:"enabled"`
	Variations        []EcomMtCatalogProductStockVariationDTO `json:"variations"`
}
