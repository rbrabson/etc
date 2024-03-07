package ftc

import (
	"encoding/json"
	"strings"
	"time"
)

type FtcTime time.Time

const (
	dateFmt = "2006-01-02T15:04:05"
)

// UnmarshalJSON parses the json time into a time value
func (ft *FtcTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateFmt, s)
	if err != nil {
		return err
	}
	*ft = FtcTime(t)
	return nil
}

// MarshalJSON returns a JSON encoding of the time
func (ft FtcTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ft))
}

// Format function for printing the date
func (ft FtcTime) Format() string {
	t := time.Time(ft)
	return t.Format(dateFmt)
}

// String function for printing the date
func (ft FtcTime) String() string {
	return ft.Format()
}
