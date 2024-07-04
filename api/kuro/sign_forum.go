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

type GetForumTaskData struct {
	CurrentDailyGold int              `json:"currentDailyGold"`
	MaxDailyGold     int              `json:"maxDailyGold"`
	GrowTask         []*ForumTaskInfo `json:"growTask"`
	DailyTask        []*ForumTaskInfo `json:"dailyTask"`
}

type ForumTaskInfo struct {
	CompleteTimes   int     `json:"completeTimes"`
	GainGold        int     `json:"gainGold"`
	NeedActionTimes int     `json:"needActionTimes"`
	Process         float64 `json:"process"`
	Remark          string  `json:"remark"`
	SkipType        int     `json:"skipType"`
	Times           int     `json:"times"`
}

func GetForumTask(account config.Account) (*GetForumTaskData, error) {
	req := R(account).SetFormData(gh.MS{"userId": "", "gameId": "0"})
	return Exec[*GetForumTaskData](req, "POST", "/encourage/level/getTaskProcess")
}

type GetForumGoldData struct {
	GoldNum int `json:"goldNum"`
}

func GetForumGold(account config.Account) (*GetForumGoldData, error) {
	req := R(account).SetFormData(gh.MS{"userId": ""})
	return Exec[*GetForumGoldData](req, "POST", "/encourage/gold/getTotalGold")
}
