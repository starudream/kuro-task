package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/kuro-task/config"
)

func TestGetUser(t *testing.T) {
	data, err := GetUser(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
