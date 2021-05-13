package catalog

type BaseProductID struct {
	ID            int64  `json:"id" bson:"id"`
	AlternativeID *int64 `json:"alternative_id" bson:"alternative_id"`
}
type BaseProductInfo struct {
	SKU               string   `json:"sku" bson:"sku"`
	Name              string   `json:"name" bson:"name"`
	Brand             string   `json:"brand" bson:"brand"`
	CustomReferenceID *string  `json:"custom_reference_id" bson:"custom_reference_id"`
	DefaultImageURL   *string  `json:"default_image_url" bson:"default_image_url"`
	DescriptionText   *string  `json:"description_text" bson:"description_text"`
	DescriptionHTML   *string  `json:"description_html" bson:"description_html"`
	ShortSummary      *string  `json:"short_summary" bson:"short_summary"`
	Enabled           bool     `json:"enabled" bson:"enabled"`
	Images            []string `json:"images" bson:"images"`
}

type BaseProductVariation struct {
	Type  string `json:"type" bson:"type"`
	Label string `json:"label" bson:"label"`
	Value string `json:"value" bson:"value"`
}
type PublicProduct struct {
	BaseProductID
	BaseProductInfo
	Categories []string               `json:"categories"`
	Variations []BaseProductVariation `json:"variations"`
}

type CollectionProductCategory struct {
	EcomID int64
	CUID   string
	Name   string
}

type CollectionProductVariation struct {
	BaseProductVariation `bson:",inline"`
	AlternativeID        *int64 `json:"alternative_id" bson:"alternative_id"`
}

type CollectionProduct struct {
	BaseProductID   `bson:",inline"`
	BaseProductInfo `bson:",inline"`
	RetailerID      int64                        `json:"retailer_id" bson:"retailer_id"`
	Categories      []CollectionProductCategory  `json:"categories" bson:"categories"`
	Variations      []CollectionProductVariation `json:"variations" bson:"variations"`
}

type EcomMtCatalogProduct struct {
	BaseProductInfo
	CategoryID *int64                 `json:"category_id"`
	Variations []BaseProductVariation `json:"variations"`
}
