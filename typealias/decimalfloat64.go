package typealias

import (
	"encoding/json"
	"fmt"
	"math/big"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DecimalFloat64 big.Rat

func (bd *DecimalFloat64) UnmarshalJSON(b []byte) error {
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}

	br := new(big.Rat)
	br.SetFloat64(f)

	*bd = DecimalFloat64(*br)

	return nil
}

func (bd DecimalFloat64) MarshalJSON() ([]byte, error) {
	f, _ := bd.Float64()

	return json.Marshal(f)
}

func NewDecimalFloat64F(value float64) *DecimalFloat64 {
	br := big.Rat{}
	br.SetFloat64(value)

	bd := DecimalFloat64(br)
	return &bd
}

func NewDecimalFloat64S(value string) (*DecimalFloat64, error) {
	br := big.Rat{}
	_, ok := br.SetString(value)

	if !ok {
		return nil, fmt.Errorf("Error parsing DecimalFloat from %s", value)
	}

	bd := DecimalFloat64(br)
	return &bd, nil
}

func NewDecimalFloat64D(value primitive.Decimal128) (*DecimalFloat64, error) {
	return NewDecimalFloat64S(value.String())
}

func (bd DecimalFloat64) FloatString() string {
	br := big.Rat(bd)

	return br.FloatString(DecimalPrecision)
}

func (bd DecimalFloat64) Float64() (float64, bool) {
	br := big.Rat(bd)

	return br.Float64()
}

func (bd DecimalFloat64) Float64L() float64 {
	br := big.Rat(bd)

	f, _ := br.Float64()
	return f
}

func (bd DecimalFloat64) BsonDecimal128() primitive.Decimal128 {
	bdCore := Decimal(bd)
	return bdCore.BsonDecimal128()
}
