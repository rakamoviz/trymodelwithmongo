package typealias

import (
	"math/big"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	BigDecimalPrecision = 4
)

type BigDecimal big.Rat

func (bd BigDecimal) BsonDecimal128() primitive.Decimal128 {
	br := big.Rat(bd)
	d, _ := primitive.ParseDecimal128(br.FloatString(BigDecimalPrecision))

	return d
}
