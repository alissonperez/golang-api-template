package resources

import (
	"fmt"
	"strings"
	"time"
)

const ISO8601Layout = "2006-01-02T15:04:05.999-07:00"

var nilTime = (time.Time{}).UnixNano()

func DateTime8601FromString(dateTimeStr string) (DateTime8601, error) {
	dateTime := DateTime8601{}
	err := dateTime.FromString(dateTimeStr)
	return dateTime, err
}

type DateTime8601 struct {
	time.Time
}

func (dt *DateTime8601) FromString(dateTimeStr string) (err error) {
	dt.Time, err = time.Parse(ISO8601Layout, dateTimeStr)
	return
}

func (dt *DateTime8601) UnmarshalJSON(b []byte) (err error) {
	return dt.FromString(strings.Trim(string(b), "\""))
}

func (dt *DateTime8601) MarshalJSON() ([]byte, error) {
	if dt.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%q", dt.Time.Format(ISO8601Layout))), nil
}
