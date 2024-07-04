package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestGetRoleBase(t *testing.T) {
	role, account := GetFirstRole(t)
	data, err := GetRoleBase(role.GameId, role.ServerId, role.RoleId, account)
	testutil.LogNoErr(t, err, data)
}

func TestGetRoleList(t *testing.T) {
	role, account := GetFirstRole(t)
	data, err := GetRoleList(role.GameId, role.ServerId, role.RoleId, account)
	testutil.LogNoErr(t, err, data)
}

func TestGetRoleCalabash(t *testing.T) {
	role, account := GetFirstRole(t)
	data, err := GetRoleCalabash(role.GameId, role.ServerId, role.RoleId, account)
	testutil.LogNoErr(t, err, data)
}

func TestGetRoleChallenge(t *testing.T) {
	role, account := GetFirstRole(t)
	data, err := GetRoleChallenge(role.GameId, role.ServerId, role.RoleId, 19, GameMCCountryCodeHL, account)
	testutil.LogNoErr(t, err, data)
}

func TestGetRoleExplore(t *testing.T) {
	role, account := GetFirstRole(t)
	data, err := GetRoleExplore(role.GameId, role.ServerId, role.RoleId, 19, GameMCCountryCodeHL, account)
	testutil.LogNoErr(t, err, data)
}
