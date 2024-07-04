package kuro

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Timestamp struct {
	time.Time
}

var _ json.Unmarshaler = (*Timestamp)(nil)

func (t *Timestamp) UnmarshalJSON(bs []byte) error {
	s := strings.Trim(string(bs), `"`)
	i, _ := strconv.Atoi(s)
	switch len(s) {
	case 10:
		t.Time = time.Unix(int64(i), 0)
	case 13:
		t.Time = time.Unix(0, int64(i)*1e6)
	}
	return nil
}
