package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/kuro-task/config"
	"github.com/starudream/kuro-task/util"
)

func TestSendPhoneCodeGeeTest(t *testing.T) {
	data, err := SendPhoneCodeGeeTest(config.C().FirstAccount().Phone, nil)
	testutil.LogNoErr(t, err, data)
}

func TestLoginByPhoneCode(t *testing.T) {
	data, err := LoginByPhoneCode(config.C().FirstAccount().Phone, "123456", util.RandString(util.CharsetHex, 40))
	testutil.LogNoErr(t, err, data)
}
