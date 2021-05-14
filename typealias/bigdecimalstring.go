package typealias

import (
	"encoding/json"
	"fmt"
	"math/big"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BigDecimalString big.Rat

func (bd *BigDecimalString) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	br := new(big.Rat)
	_, ok := br.SetString(s)

	if !ok {
		return fmt.Errorf("Cannot parse %s to BigDecimalString", s)
	}

	*bd = BigDecimalString(*br)

	return nil
}

func (bd BigDecimalString) MarshalJSON() ([]byte, error) {
	return json.Marshal(bd.FloatString())
}

func NewBigDecimalStringF(value float64) *BigDecimalString {
	br := big.Rat{}
	br.SetFloat64(value)

	bd := BigDecimalString(br)
	return &bd
}

func NewBigDecimalStringS(value string) (*BigDecimalString, error) {
	br := big.Rat{}
	_, ok := br.SetString(value)

	if !ok {
		return nil, fmt.Errorf("Error parsing BigDecimalString from %s", value)
	}

	bd := BigDecimalString(br)
	return &bd, nil
}

func (bd BigDecimalString) FloatString() string {
	br := big.Rat(bd)

	return br.FloatString(BigDecimalPrecision)
}

func (bd BigDecimalString) Float64() (float64, bool) {
	br := big.Rat(bd)

	return br.Float64()
}

func (bd BigDecimalString) BsonDecimal128() primitive.Decimal128 {
	bdCore := BigDecimal(bd)
	return bdCore.BsonDecimal128()
}
