package catalog

type RetailerProduct struct {
	ID              int64                  `json:"id"`
	Name            string                 `json:"name"`
	ParentID        *int64                 `json:"parentId"`
	Description     string                 `json:"longDescription"`
	Images          []ProductStoreImage    `json:"images"`
	Status          string                 `json:"status"`
	MasterProductID int64                  `json:"productId"`
	Attributes      map[string]interface{} `json:"attributes"`
	Retailer        struct {
		ID int64 `json:"id"`
	} `json:"retailer"`
	SellType struct {
		MaxQuantity int `json:"maxQuantity"`
	} `json:"sellType"`
}
