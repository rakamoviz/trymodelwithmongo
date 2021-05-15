package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/httphandler"
)

type HttpHandler struct {
	httphandler.Base
}

func Setup(
	baseGroup *echo.Group, httpHandler HttpHandler, path string,
) {
	routes := baseGroup.Group(path)
	routes.POST("/productStock/:retailerID", httpHandler.PostProductStock)
}
