package util

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestRandString(t *testing.T) {
	testutil.Log(t, RandString("0123456789ABCDEF", 40))
}
