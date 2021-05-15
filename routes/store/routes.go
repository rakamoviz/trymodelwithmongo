package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/routes"
)

type Route struct {
	routes.Route
}

func Setup(
	server *echo.Echo, baseRoute routes.Route,
) {
	storeRoute := &Route{baseRoute}

	group := server.Group("/store")
	group.POST("/productStock/:retailerID", storeRoute.PostProductStock)
}
