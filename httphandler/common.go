package httphandler

import (
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
)

type Base struct {
	DB               drivers.DB
	EcomMtCatalogAPI *ecomMtCatalogAPI.EcomMtCatalogAPI
}
