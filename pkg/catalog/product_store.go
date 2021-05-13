package catalog

import (
	"fmt"
	"time"
)

const StatusPublished = "published"

type ProductStoreImage struct {
	Position int    `json:"position"`
	Path     string `json:"path"`
}

type ProductStore struct {
	// basic properties
	VirtualProductID int64                  `json:"id"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Images           []ProductStoreImage    `json:"images"`
	SKU              *string                `json:"sku"`
	EAN              *string                `json:"ean"`
	Stock            int                    `json:"stock"`
	Price            float64                `json:"price"`
	InStock          bool                   `json:"in_stock"`
	IsAvailable      bool                   `json:"is_available"`
	RetailerMarkdown float64                `json:"retailer_markdown"`
	BrandMarkdown    float64                `json:"brand_markdown"`
	Attributes       map[string]interface{} `json:"attributes"`

	// catalog properties
	VirtualStoreID        int64  `json:"store_id"`
	RetailerProductID     int64  `json:"retailer_product_id"`
	MasterProductID       int64  `json:"master_product_id"`
	RetailerID            int64  `json:"retailer_id"`
	MasterParentProductID *int64 `json:"master_parent_product_id"`

	RetailerProduct RetailerProduct `json:"-"`
	VirtualProduct  VirtualProduct  `json:"-"`

	// catalog-clg properties
	CreatedAt       time.Time `json:"-"`
	Key             string    `json:"-"`
	VariationFamily string    `json:"-" bson:"variationFamily"`
}

func NewProductStore(rp RetailerProduct, vp VirtualProduct) ProductStore {
	vf := fmt.Sprintf("%d_%d", rp.MasterProductID, vp.VirtualStoreID)
	if rp.ParentID != nil {
		vf = fmt.Sprintf("%d_%d", *rp.ParentID, vp.VirtualStoreID)
	}

	return ProductStore{
		VirtualProductID: vp.ID,
		Name:             rp.Name,
		Description:      rp.Description,
		Images:           rp.Images,
		SKU:              &vp.SKU,
		EAN:              &vp.EAN,
		Stock:            vp.Stock,
		Price:            vp.VPrice,
		InStock:          vp.InStock,
		IsAvailable:      vp.Status == StatusPublished && rp.Status == StatusPublished,
		RetailerMarkdown: vp.RetailerMarkdown,
		BrandMarkdown:    vp.BrandMarkdown,
		Attributes:       rp.Attributes,

		VirtualStoreID:        vp.VirtualStoreID,
		RetailerProductID:     rp.ID,
		MasterProductID:       rp.MasterProductID,
		RetailerID:            rp.Retailer.ID,
		MasterParentProductID: rp.ParentID,

		RetailerProduct: rp,
		VirtualProduct:  vp,

		Key:             fmt.Sprintf("%d_%d", vp.ID, vp.VirtualStoreID),
		VariationFamily: vf,
	}
}
