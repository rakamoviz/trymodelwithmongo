package main

import (
	"time"

	"bitbucket.org/rappinc/gohttp"
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
	routesStore "github.com/rakamoviz/trymodelwithmongo/routes/store"
	"github.com/rakamoviz/trymodelwithmongo/util"
)

func main() {
	db, err := drivers.GetDriver("mongodb://localhost:27017", "ecomcatalogclg")
	if err != nil {
		panic(err)
	}

	ecomMtCatalogAPIClient := ecomMtCatalogAPI.New(gohttp.NewClient(gohttp.Options{
		ClientName: "ecom-catalog-clg",
		Timeout:    30 * time.Second,
	}), ecomMtCatalogAPI.Option{
		Host:                "",
		Country:             "dev",
		EcomMtCatalogAPIKey: "",
		ClientName:          "ecom-catalog-clg",
	})

	echoServer := echo.New()
	commonShell := util.CommonShell{
		DB:               db,
		EcomMtCatalogAPI: ecomMtCatalogAPIClient,
	}

	routesStore.Setup(echoServer, commonShell)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))
}
