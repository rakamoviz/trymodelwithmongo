package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
)

type Shell struct {
	DB               drivers.DB
	EcomMtCatalogAPI *ecomMtCatalogAPI.EcomMtCatalogAPI
}

func Setup(
	server *echo.Echo, DB drivers.DB,
	EcomMtCatalogAPI *ecomMtCatalogAPI.EcomMtCatalogAPI,
) {
	shell := &Shell{DB, EcomMtCatalogAPI}

	group := server.Group("/store")
	group.POST("/productStock/:retailerID", shell.PostProductStock)
}
