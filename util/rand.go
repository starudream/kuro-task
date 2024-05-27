package util

import (
	"math/rand/v2"
	"strings"
)

const CharsetHex = "0123456789ABCDEF"

func RandString(charset string, length int) string {
	ml := uint32(len(charset))
	sb := strings.Builder{}
	for i := 1; i < length; i++ {
		sb.WriteByte(charset[rand.Uint32N(ml)])
	}
	return sb.String()
}
