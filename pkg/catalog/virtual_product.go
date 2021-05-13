package catalog

type VirtualProduct struct {
	ID                int64   `json:"id"`
	RetailerProductID int64   `json:"retailerProductId"`
	ProductID         int64   `json:"productId"`
	StoreType         string  `json:"storeType"`
	Country           string  `json:"country"`
	SKU               string  `json:"sku"`
	EAN               string  `json:"ean"`
	Stock             int     `json:"stock"`
	VPrice            float64 `json:"price"`
	RetailerMarkdown  float64 `json:"retailerMarkdown"`
	BrandMarkdown     float64 `json:"brandMarkdown"`
	InStock           bool    `json:"inStock"`
	BalancePrice      float64 `json:"balancePrice"`
	VirtualStoreID    int64   `json:"virtualStoreId"`
	SecurityStock     int     `json:"securityStock"`
	Status            string  `json:"status"`
}
