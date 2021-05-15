package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/httphandler"
)

type HttpHandler struct {
	httphandler.Base
}

func Setup(
	baseGroup *echo.Group, baseHttpHandler httphandler.Base,
) {
	httpHandler := &HttpHandler{baseHttpHandler}

	group := baseGroup.Group("/store")
	group.POST("/productStock/:retailerID", httpHandler.PostProductStock)
}
