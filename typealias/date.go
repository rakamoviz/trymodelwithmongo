package typealias

import (
	"strings"
	"time"
)

type Date time.Time

const ctLayout = "2006-11-02"

func (ct *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = Date(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(ct).String()), nil
}
