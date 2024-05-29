package kuro

import (
	"strconv"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/kuro-task/config"
)

type SignForumData struct {
	GainVoList   []*GainVo `json:"gainVoList"`
	ContinueDays int       `json:"continueDays"`
}

type GainVo struct {
	GainTyp   int `json:"gainTyp"`
	GainValue int `json:"gainValue"`
}

func SignForum(gid int, account config.Account) (*SignForumData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid)})
	return Exec[*SignForumData](req, "POST", "/user/signIn")
}

type GetSignForumData struct {
	ContinueDays int  `json:"continueDays"`
	HasSignIn    bool `json:"hasSignIn"`
}

func GetSignForum(gid int, account config.Account) (*GetSignForumData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid)})
	return Exec[*GetSignForumData](req, "POST", "/user/signIn/info")
}
