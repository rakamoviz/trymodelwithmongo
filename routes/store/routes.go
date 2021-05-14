package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/util"
)

type RouteEnv struct {
	util.CommonRouteEnv
}

func Setup(
	server *echo.Echo, commonRouteEnv util.CommonRouteEnv,
) {
	routeEnv := &RouteEnv{commonRouteEnv}

	group := server.Group("/store")
	group.POST("/productStock/:retailerID", routeEnv.PostProductStock)
}
