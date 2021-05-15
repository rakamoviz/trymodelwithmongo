package routes

import (
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
)

type Route struct {
	DB               drivers.DB
	EcomMtCatalogAPI *ecomMtCatalogAPI.EcomMtCatalogAPI
}
