package api

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/rappinc/gohttp"
	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
)

type Option struct {
	Host                string
	Country             string
	EcomMtCatalogAPIKey string
	ClientName          string
}

type EcomMtCatalogAPI struct {
	host       string
	country    string
	clientName string
	client     gohttp.Client

	ecomMtCatalogAPIKey string
}

func New(client gohttp.Client, opts Option) *EcomMtCatalogAPI {
	log.Printf("[ecom-mt-catalog/api] Host: %s\n", opts.Host)

	return &EcomMtCatalogAPI{
		host:       opts.Host,
		country:    opts.Country,
		clientName: opts.ClientName,
		client:     client,

		ecomMtCatalogAPIKey: opts.EcomMtCatalogAPIKey,
	}
}

func (api *EcomMtCatalogAPI) headers() map[string]string {
	return map[string]string{
		"api_key":       api.ecomMtCatalogAPIKey,
		"Authorization": api.ecomMtCatalogAPIKey,
		"x-user-id":     "100",
		"x-user-name":   api.clientName,
		"x-country":     api.country,
	}
}

func (api *EcomMtCatalogAPI) SyncProductStock(productStock catalog.EcomMtCatalogProductStock) error {
	req := gohttp.Req{
		Method:       http.MethodPut,
		Path:         fmt.Sprintf("/api/ecom-mt-catalog/cpgs/catalog/store/%d/product/stock?total=true", productStock.StoreID),
		ExtraHeaders: api.headers(),
		Tx:           nil,
		Host:         api.host,
	}

	res := gohttp.JSONRequestWithResult(api.client, req, productStock, nil)

	if res.Error != nil {
		return fmt.Errorf("error in UpdatePublicProduct Ecom-API - error: %v", res.Error)
	} else {
		if res.ResponseCode != http.StatusOK {
			return fmt.Errorf("error in UpdatePublicProduct Ecom-API - status: %d", res.ResponseCode)
		}
	}

	return nil
}
