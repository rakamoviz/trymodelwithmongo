package store

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
	"github.com/rakamoviz/trymodelwithmongo/util"
)

func (shell *Shell) PostProductStock(c echo.Context) error {
	retailerID, err := util.StringToInt64(c.Param("retailerID"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	publicProductStock := new(catalog.PublicProductStock)
	if err = c.Bind(publicProductStock); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	collectionProductFromDB, err := shell.DB.FindPublicProductBySKUAndRetailerID(
		publicProductStock.ProductSKU,
		retailerID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if collectionProductFromDB == nil {
		return c.JSON(http.StatusNotFound, fmt.Sprintf(
			"Product not found for sku %s and retailerID %d",
			publicProductStock.ProductSKU, retailerID,
		))
	}

	collectionProductStockVariations := []catalog.CollectionProductStockVariation{}
	for _, publicProductStockVariation := range publicProductStock.Variations {
		collectionProductStockVariations = append(
			collectionProductStockVariations, catalog.CollectionProductStockVariation{
				VariationValue: publicProductStockVariation.VariationValue,
				Stock:          publicProductStockVariation.Stock,
				PriceDelta:     publicProductStockVariation.PriceDelta.BsonDecimal128(),
			},
		)
	}

	fmt.Println("publicProductStock", publicProductStock.UnitPrice.FloatString())
	collectionProductStock := &catalog.CollectionProductStock{
		StoreID:           publicProductStock.StoreID,
		ProductSKU:        publicProductStock.ProductSKU,
		RetailerID:        retailerID,
		RetailerProductID: collectionProductFromDB.AlternativeID,
		Stock:             publicProductStock.Stock,
		UnitPrice:         publicProductStock.UnitPrice.BsonDecimal128(),
		Enabled:           publicProductStock.Enabled,
		Variations:        collectionProductStockVariations,
	}

	return c.JSON(http.StatusOK, collectionProductStock)
}
