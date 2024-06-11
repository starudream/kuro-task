package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/kuro-task/config"
)

func TestSendPhoneCodeGeeTest(t *testing.T) {
	data, err := SendPhoneCodeGeeTest(config.C().FirstAccount().Phone, nil, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestLoginByPhoneCode(t *testing.T) {
	data, err := LoginByPhoneCode(config.C().FirstAccount().Phone, "123456", config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
