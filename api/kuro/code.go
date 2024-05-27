package kuro

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Code int

var _ json.Unmarshaler = (*Code)(nil)

func (c *Code) UnmarshalJSON(bs []byte) error {
	i, _ := strconv.Atoi(strings.Trim(string(bs), `"`))
	*c = Code(i)
	return nil
}
