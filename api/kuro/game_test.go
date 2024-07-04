package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/kuro-task/config"
)

func TestGetGame3Widget(t *testing.T) {
	data, err := GetGame3Widget(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestListRole(t *testing.T) {
	data, err := ListRole(GameIdMC, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func GetFirstRole(t *testing.T) (*Role, config.Account) {
	account := config.C().FirstAccount()
	data, err := ListRole(GameIdMC, account)
	testutil.LogNoErr(t, err, data)
	return data[0], account
}
