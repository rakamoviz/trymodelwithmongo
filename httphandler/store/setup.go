package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/httphandler"
)

type HttpHandler struct {
	httphandler.Base
}

func Setup(
	server *echo.Echo, baseHttpHandler httphandler.Base,
) {
	httpHandler := &HttpHandler{baseHttpHandler}

	group := server.Group("/store")
	group.POST("/productStock/:retailerID", httpHandler.PostProductStock)
}
