package ftc

import (
	"encoding/json"
	"strings"
	"time"
)

type Time time.Time

const (
	dateFmt = "2006-01-02T15:04:05"
)

// UnmarshalJSON parses the json time into a time value
func (ft *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateFmt, s)
	if err != nil {
		return err
	}
	*ft = Time(t)
	return nil
}

// MarshalJSON returns a JSON encoding of the time
func (ft Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ft))
}

// Format function for printing the date
func (ft Time) Format() string {
	t := time.Time(ft)
	str := t.Format(dateFmt)
	str = strings.TrimSuffix(str, "Z")
	return str
}

// String function for printing the date
func (ft Time) String() string {
	return ft.Format()
}
