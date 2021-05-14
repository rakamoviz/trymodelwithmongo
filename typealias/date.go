package typealias

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

const layoutISO = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(layoutISO, s)
	*d = Date(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(d).Format(layoutISO))), nil
}
