package kuro

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/kuro-task/config"
)

func TestListPost(t *testing.T) {
	data, err := ListPost(GameIdMC, ForumIdMC10, 1, config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestGetPost(t *testing.T) {
	data, err := GetPost("1256182991485353984", config.C().FirstAccount())
	testutil.LogNoErr(t, err, data)
}

func TestLikePost(t *testing.T) {
	err := LikePost(GameIdMC, ForumIdMC10, "1256182991485353984", "10381395", false, config.C().FirstAccount())
	testutil.LogNoErr(t, err)
}

func TestSharePost(t *testing.T) {
	err := SharePost(GameIdMC, config.C().FirstAccount())
	testutil.LogNoErr(t, err)
}
