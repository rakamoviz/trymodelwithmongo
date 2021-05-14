package util

import (
	"strconv"

	"github.com/rakamoviz/trymodelwithmongo/pkg/drivers"
	ecomMtCatalogAPI "github.com/rakamoviz/trymodelwithmongo/pkg/ecom-mt-catalog/api"
)

func StringToInt64(value string) (int64, error) {
	int64Val, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return int64(int64Val), nil
}

type CommonRouteEnv struct {
	DB               drivers.DB
	EcomMtCatalogAPI *ecomMtCatalogAPI.EcomMtCatalogAPI
}
