package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/kuro-task/config"
)

func TestSignGame(t *testing.T) {
	role := GetFirstRole(t)
	data, err := SignGame(role.GameId, role.ServerId, role.RoleId, role.UserId, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestListSignGame(t *testing.T) {
	role := GetFirstRole(t)
	data, err := ListSignGame(role.GameId, role.ServerId, role.RoleId, role.UserId, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, data.GoodsMap())
}

func TestListSignGameRecord(t *testing.T) {
	role := GetFirstRole(t)
	data, err := ListSignGameRecord(role.GameId, role.ServerId, role.RoleId, role.UserId, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data, len(data), data.Today(), data.Today().ShortString())
}
