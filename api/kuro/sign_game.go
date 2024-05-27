package kuro

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/kuro-task/config"
)

type Good struct {
	SerialNum int    `json:"serialNum,omitempty"`
	Type      int    `json:"type,omitempty"`
	GoodsName string `json:"goodsName"`
	GoodsNum  int    `json:"goodsNum"`
	GoodsId   int    `json:"goodsId"`
	SigInDate string `json:"sigInDate,omitempty"`
}

type Goods []*Good

func (v1 Goods) Today() (v2 Goods) {
	today := time.Now().Format(time.DateOnly)
	for i := range v1 {
		if strings.HasPrefix(v1[i].SigInDate, today) {
			v2 = append(v2, v1[i])
		}
	}
	return
}

func (v1 Goods) ShortString() string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = v.GoodsName + "*" + strconv.Itoa(v.GoodsNum)
	}
	return strings.Join(v2, ", ")
}

func (v1 Goods) ShortStringByMap(m map[int]*Good) string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = m[v.GoodsId].GoodsName + "*" + strconv.Itoa(v.GoodsNum)
	}
	return strings.Join(v2, ", ")
}

type SignGameData struct {
	TodayList    Goods `json:"todayList"`
	TomorrowList Goods `json:"tomorrowList"`
}

func SignGame(gid int, sid, rid string, uid int, account config.Account) (*SignGameData, error) {
	month := fmt.Sprintf("%02d", time.Now().Month())
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "userId": strconv.Itoa(uid), "reqMonth": month})
	return Exec[*SignGameData](req, "POST", "/encourage/signIn/v2")
}

type ListSignGameData struct {
	DisposableGoodsList Goods `json:"disposableGoodsList"`
	DisposableSignNum   int   `json:"disposableSignNum"`

	SignInGoodsConfigs Goods `json:"signInGoodsConfigs"`
	SignLoopGoodsList  Goods `json:"signLoopGoodsList"`
	SigInNum           int   `json:"sigInNum"`

	NowServerTimes  string `json:"nowServerTimes"`
	EventStartTimes string `json:"eventStartTimes"`
	EventEndTimes   string `json:"eventEndTimes"`
	ExpendGold      int    `json:"expendGold"`
	ExpendNum       int    `json:"expendNum"`
	IsSigIn         bool   `json:"isSigIn"`
	OmissionNnm     int    `json:"omissionNnm"`
}

func (v *ListSignGameData) GoodsMap() map[int]*Good {
	m := map[int]*Good{}
	for _, good := range v.DisposableGoodsList {
		m[good.GoodsId] = good
	}
	for _, good := range v.SignInGoodsConfigs {
		m[good.GoodsId] = good
	}
	for _, good := range v.SignLoopGoodsList {
		m[good.GoodsId] = good
	}
	return m
}

func ListSignGame(gid int, sid, rid string, uid int, account config.Account) (*ListSignGameData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "userId": strconv.Itoa(uid)})
	return Exec[*ListSignGameData](req, "POST", "/encourage/signIn/initSignInV2")
}

func ListSignGameRecord(gid int, sid, rid string, uid int, account config.Account) (Goods, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "userId": strconv.Itoa(uid)})
	return Exec[Goods](req, "POST", "/encourage/signIn/queryRecordV2")
}
