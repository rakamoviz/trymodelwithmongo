package typealias

import (
	"math/big"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DecimalPrecision = 4
)

type Decimal big.Rat

func (bd Decimal) BsonDecimal128() primitive.Decimal128 {
	br := big.Rat(bd)
	d, _ := primitive.ParseDecimal128(br.FloatString(DecimalPrecision))

	return d
}
