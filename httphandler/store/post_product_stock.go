package store

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rakamoviz/trymodelwithmongo/pkg/catalog"
	"github.com/rakamoviz/trymodelwithmongo/typealias"
	"github.com/rakamoviz/trymodelwithmongo/util"
)

func toMTCatalogProductStock(
	ctx context.Context, ps catalog.ProductStockEntity,
) (catalog.EcomMtCatalogProductStockDTO, error) {
	unitPrice, err := typealias.NewDecimalFloat64D(ps.UnitPrice)
	if err != nil {
		return catalog.EcomMtCatalogProductStockDTO{}, err
	}

	mtCatalogPS := catalog.EcomMtCatalogProductStockDTO{
		StoreID:           ps.StoreID,
		RetailerProductID: ps.RetailerProductID,
		Stock:             ps.Stock,
		UnitPrice:         unitPrice,
		Enabled:           ps.Enabled,
	}

	mtCatalogProductStockVariations := []catalog.EcomMtCatalogProductStockVariationDTO{}
	for _, psVariation := range ps.Variations {
		priceDelta, err := typealias.NewDecimalFloat64D(psVariation.PriceDelta)
		if err != nil {
			return mtCatalogPS, err
		}

		mtCatalogProductStockVariations = append(
			mtCatalogProductStockVariations, catalog.EcomMtCatalogProductStockVariationDTO{
				VariationValue: psVariation.VariationValue,
				Stock:          psVariation.Stock,
				PriceDelta:     priceDelta,
			},
		)
	}

	return mtCatalogPS, nil
}

func (handler *HttpHandler) PostProductStock(c echo.Context) error {
	retailerID, err := util.StringToInt64(c.Param("retailerID"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	productStockDTO := new(catalog.ProductStockDTO)
	if err = c.Bind(productStockDTO); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	productEntity, err := handler.DB.FindPublicProductBySKUAndRetailerID(
		productStockDTO.ProductSKU,
		retailerID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if productEntity == nil {
		return c.JSON(http.StatusNotFound, fmt.Sprintf(
			"Product not found for sku %s and retailerID %d",
			productStockDTO.ProductSKU, retailerID,
		))
	}

	productStockVariationEntities := []catalog.ProductStockVariationEntity{}
	for _, productStockVariationDTO := range productStockDTO.Variations {
		productStockVariationEntities = append(
			productStockVariationEntities, catalog.ProductStockVariationEntity{
				VariationValue: productStockVariationDTO.VariationValue,
				Stock:          productStockVariationDTO.Stock,
				PriceDelta:     productStockVariationDTO.PriceDelta.BsonDecimal128(),
			},
		)
	}

	productStockEntity := catalog.ProductStockEntity{
		ProductStockEntityKey: catalog.ProductStockEntityKey{
			StoreID:    productStockDTO.StoreID,
			ProductSKU: productStockDTO.ProductSKU,
			RetailerID: retailerID,
		},
		RetailerProductID: productEntity.AlternativeID,
		Stock:             productStockDTO.Stock,
		UnitPrice:         productStockDTO.UnitPrice.BsonDecimal128(),
		Enabled:           productStockDTO.Enabled,
		Variations:        productStockVariationEntities,
	}

	err = handler.DB.SaveProductStock(productStockEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	savedProductStockEntity, err := handler.DB.FindProductStock(productStockEntity.ProductStockEntityKey)

	mtCatalogProductStockDTO, err := toMTCatalogProductStock(c.Request().Context(), savedProductStockEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = handler.EcomMtCatalogAPI.SyncProductStock(mtCatalogProductStockDTO)
	if err != nil {
		//return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, savedProductStockEntity)
}
