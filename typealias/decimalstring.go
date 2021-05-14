package typealias

import (
	"encoding/json"
	"fmt"
	"math/big"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DecimalString big.Rat

func (bd *DecimalString) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	br := new(big.Rat)
	_, ok := br.SetString(s)

	if !ok {
		return fmt.Errorf("Cannot parse %s to DecimalString", s)
	}

	*bd = DecimalString(*br)

	return nil
}

func (bd DecimalString) MarshalJSON() ([]byte, error) {
	return json.Marshal(bd.FloatString())
}

func NewDecimalStringF(value float64) *DecimalString {
	br := big.Rat{}
	br.SetFloat64(value)

	bd := DecimalString(br)
	return &bd
}

func NewDecimalStringS(value string) (DecimalString, error) {
	br := big.Rat{}
	_, ok := br.SetString(value)

	if !ok {
		return DecimalString{}, fmt.Errorf("Error parsing DecimalString from %s", value)
	}

	return DecimalString(br), nil
}

func (bd DecimalString) FloatString() string {
	br := big.Rat(bd)

	return br.FloatString(DecimalPrecision)
}

func (bd DecimalString) Float64() (float64, bool) {
	br := big.Rat(bd)

	return br.Float64()
}

func (bd DecimalString) BsonDecimal128() primitive.Decimal128 {
	bdCore := Decimal(bd)
	return bdCore.BsonDecimal128()
}
