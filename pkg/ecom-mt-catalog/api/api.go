package api

import (
	"log"

	"bitbucket.org/rappinc/gohttp"
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
