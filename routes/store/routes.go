package store

import (
	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/util"
)

type Shell struct {
	util.CommonShell
}

func Setup(
	server *echo.Echo, commonShell util.CommonShell,
) {
	shell := &Shell{commonShell}

	group := server.Group("/store")
	group.POST("/productStock/:retailerID", shell.PostProductStock)
}
