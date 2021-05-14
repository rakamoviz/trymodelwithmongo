package typealias

import (
	"encoding/json"
	"fmt"
	"math/big"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BigDecimalFloat64 big.Rat

func (bd *BigDecimalFloat64) UnmarshalJSON(b []byte) error {
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}

	br := new(big.Rat)
	br.SetFloat64(f)

	*bd = BigDecimalFloat64(*br)

	return nil
}

func (bd BigDecimalFloat64) MarshalJSON() ([]byte, error) {
	f, _ := bd.Float64()

	return json.Marshal(f)
}

func NewBigDecimalFloat64F(value float64) *BigDecimalFloat64 {
	br := big.Rat{}
	br.SetFloat64(value)

	bd := BigDecimalFloat64(br)
	return &bd
}

func NewBigDecimalFloat64S(value string) (*BigDecimalFloat64, error) {
	br := big.Rat{}
	_, ok := br.SetString(value)

	if !ok {
		return nil, fmt.Errorf("Error parsing BigDecimalFloat from %s", value)
	}

	bd := BigDecimalFloat64(br)
	return &bd, nil
}

func (bd BigDecimalFloat64) FloatString() string {
	br := big.Rat(bd)

	return br.FloatString(BigDecimalPrecision)
}

func (bd BigDecimalFloat64) Float64() (float64, bool) {
	br := big.Rat(bd)

	return br.Float64()
}

func (bd BigDecimalFloat64) BsonDecimal128() primitive.Decimal128 {
	bdCore := BigDecimal(bd)
	return bdCore.BsonDecimal128()
}
