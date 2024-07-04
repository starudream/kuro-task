package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/kuro-task/config"
)

func TestSignForum(t *testing.T) {
	data, err := SignForum(GameIdMC, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestGetSignForum(t *testing.T) {
	data, err := GetSignForum(GameIdMC, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestGetForumTask(t *testing.T) {
	data, err := GetForumTask(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestGetForumGold(t *testing.T) {
	data, err := GetForumGold(config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}
