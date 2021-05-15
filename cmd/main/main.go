package main

import (
	"time"

	"bitbucket.org/rappinc/gohttp"
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/docdb"
	"github.com/rakamoviz/trymodelwithmongo/httphandler"
	storeHttpHandler "github.com/rakamoviz/trymodelwithmongo/httphandler/store"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
)

func main() {
	db, err := docdb.GetDriver("mongodb://localhost:27017", "ecomcatalogclg")
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

	apiRoutes := echoServer.Group("/api")

	baseHttpHandler := httphandler.Base{
		DB:               db,
		EcomMtCatalogAPI: ecomMtCatalogAPIClient,
	}

	setupStoreHttpHandler(apiRoutes, baseHttpHandler)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))
}

func setupStoreHttpHandler(baseRoutes *echo.Group, baseHttpHandler httphandler.Base) {
	httpHandler := storeHttpHandler.HttpHandler{
		Base: baseHttpHandler,
	}
	storeHttpHandler.Setup(baseRoutes, httpHandler, "/store")
}
