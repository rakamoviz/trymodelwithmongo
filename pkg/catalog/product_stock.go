package catalog

import (
	"encoding/json"
	"math/big"
	"strings"
	"time"
)

type BigDecimal big.Rat

func (bd *BigDecimal) UnmarshalJSON(b []byte) error {
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}

	br := new(big.Rat)
	br.SetFloat64(f)

	*bd = BigDecimal(*br)

	return nil
}

func (bd BigDecimal) MarshalJSON() ([]byte, error) {
	br := big.Rat(bd)

	f, _ := br.Float64()

	return json.Marshal(f)
}

type CustomTime time.Time

const ctLayout = "2006-11-02"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = CustomTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(ct).String()), nil
}

type PublicProductStockVariation struct {
	VariationValue string     `json:"variation_value"`
	Stock          int32      `json:"stock"`
	PriceDelta     BigDecimal `json:"price_delta"`
}

type PublicProductStockDiscount struct {
	DiscountType string     `json:"discount_type"`
	Value        string     `json:"value"`
	BeginDate    CustomTime `json:"begin_date"`
	EndDate      CustomTime `json:"end_date"`
}

type PublicProductStock struct {
	StoreID    int64                         `json:"store_id"`
	ProductSKU string                        `json:"product_sku"`
	Stock      int64                         `json:"stock"`
	UnitPrice  BigDecimal                    `json:"unit_price"`
	Enabled    bool                          `json:"enabled"`
	Variations []PublicProductStockVariation `json:"variations"`
	Discounts  []PublicProductStockDiscount  `json:"discounts"`
}

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

type CollectionProductStock struct {
	StoreID           int64                             `bson:"store_id"`
	ProductSKU        string                            `json:"product_sku"`
	RetailerID        int64                             `bson:"retailer_id"`
	RetailerProductID *int64                            `bson:"retailer_product_id"`
	Stock             int64                             `bson:"stock"`
	UnitPrice         BigDecimal                        `bson:"unit_price"`
	Enabled           bool                              `bson:"enabled"`
	Variations        []CollectionProductStockVariation `bson:"variations"`
	Discounts         []CollectionProductStockDiscount  `bson:"discounts"`
	CreatedAt         *time.Time                        `bson:"created_at,omitempty"`
	UpdatedAt         *time.Time                        `bson:"updated_at,omitempty"`
}

type EcomMtCatalogProductStockVariation struct {
	RetailerProductID int64      `json:"retailerProductId"`
	VariationValue    string     `json:"variationValue"`
	Stock             int32      `json:"stock"`
	PriceDelta        BigDecimal `json:"priceDelta"`
}

type EcomMtCatalogProductStock struct {
	StoreID           int64                                `json:"storeId"`
	RetailerProductID *int64                               `json:"retailerProductId"`
	Stock             int64                                `json:"stock"`
	UnitPrice         BigDecimal                           `json:"unitPrice"`
	Enabled           bool                                 `json:"enabled"`
	Variations        []EcomMtCatalogProductStockVariation `json:"variations"`
}
