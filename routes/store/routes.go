package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
	"github.com/rakamoviz/trymodelwithmongo/util"
)

type Shell struct {
	util.CommonShell
}

func Setup(
	server *echo.Echo, DB drivers.DB,
	EcomMtCatalogAPI *ecomMtCatalogAPI.EcomMtCatalogAPI,
) {
	shell := &Shell{util.CommonShell{DB, EcomMtCatalogAPI}}

	group := server.Group("/store")
	group.POST("/productStock/:retailerID", shell.PostProductStock)
}
