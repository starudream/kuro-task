package kuro

import (
	"strconv"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/kuro-task/config"
)

type GetRoleBaseData struct {
	Id               int        `json:"id"`
	Name             string     `json:"name"`
	ActiveDays       int        `json:"activeDays"`
	Level            int        `json:"level"`
	WorldLevel       int        `json:"worldLevel"`
	RoleNum          int        `json:"roleNum"`
	SoundBox         int        `json:"soundBox"`
	Energy           int        `json:"energy"`
	MaxEnergy        int        `json:"maxEnergy"`
	Liveness         int        `json:"liveness"`
	LivenessMaxCount int        `json:"livenessMaxCount"`
	LivenessUnlock   bool       `json:"livenessUnlock"`
	ChapterId        int        `json:"chapterId"`
	BigCount         int        `json:"bigCount"`
	SmallCount       int        `json:"smallCount"`
	AchievementCount int        `json:"achievementCount"`
	ShowToGuest      bool       `json:"showToGuest"`
	BoxList          []*BoxInfo `json:"boxList"`
	CreatTime        Timestamp  `json:"creatTime"`
}

type BoxInfo struct {
	Name string `json:"boxName"`
	Num  int    `json:"num"`
}

func GetRoleBase(gid int, sid, rid string, account config.Account) (*GetRoleBaseData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid})
	return Exec[*GetRoleBaseData](req, "POST", "/gamer/roleBox/aki/baseData")
}

type GetRoleListData struct {
	RoleList []*RoleInfo `json:"roleList"`
}

type RoleInfo struct {
	RoleId         int    `json:"roleId"`
	Level          int    `json:"level"`
	RoleName       string `json:"roleName"`
	StarLevel      int    `json:"starLevel"`
	AttributeId    int    `json:"attributeId"`
	AttributeName  string `json:"attributeName"`
	WeaponTypeId   int    `json:"weaponTypeId"`
	WeaponTypeName string `json:"weaponTypeName"`
	Acronym        string `json:"acronym"`
}

func GetRoleList(gid int, sid, rid string, account config.Account) (*GetRoleListData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid})
	return Exec[*GetRoleListData](req, "POST", "/gamer/roleBox/aki/roleData")
}

type GetRoleCalabashData struct {
	Level           int            `json:"level"`
	BaseCatch       string         `json:"baseCatch"`
	StrengthenCatch string         `json:"strengthenCatch"`
	CatchQuality    int            `json:"catchQuality"`
	Cost            int            `json:"cost"`
	MaxCount        int            `json:"maxCount"`
	UnlockCount     int            `json:"unlockCount"`
	PhantomList     []*PhantomInfo `json:"phantomList"`
}

type PhantomInfo struct {
	Phantom *Phantom `json:"phantom"`
	Star    int      `json:"star"`
	MaxStar int      `json:"maxStar"`
}

type Phantom struct {
	PhantomId int    `json:"phantomId"`
	Name      string `json:"name"`
	Cost      int    `json:"cost"`
	Acronym   string `json:"acronym"`
}

func GetRoleCalabash(gid int, sid, rid string, account config.Account) (*GetRoleCalabashData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid})
	return Exec[*GetRoleCalabashData](req, "POST", "/gamer/roleBox/aki/calabashData")
}

type GetRoleChallengeData struct {
	IndexList []*BossInfo `json:"indexList"`
}

type BossInfo struct {
	BossId     int    `json:"bossId"`
	BossLevel  int    `json:"bossLevel"`
	BossName   string `json:"bossName"`
	Difficulty int    `json:"difficulty"`
}

func GetRoleChallenge(gid int, sid, rid string, cid, code int, account config.Account) (*GetRoleChallengeData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "channelId": strconv.Itoa(cid), "countryCode": strconv.Itoa(code)})
	return Exec[*GetRoleChallengeData](req, "POST", "/gamer/roleBox/aki/challengeIndex")
}

type GetRoleExploreData struct {
	CountryCode       int              `json:"countryCode"`
	CountryName       string           `json:"countryName"`
	CountryProgress   string           `json:"countryProgress"`
	AreaInfoList      []*AreaInfo      `json:"areaInfoList"`
	DetectionInfoList []*DetectionInfo `json:"detectionInfoList"`
}

type AreaInfo struct {
	AreaId       int             `json:"areaId"`
	AreaProgress int             `json:"areaProgress"`
	AreaName     string          `json:"areaName"`
	ItemList     []*AreaItemInfo `json:"itemList"`
}

type AreaItemInfo struct {
	Type     int    `json:"type"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
}

type DetectionInfo struct {
	DetectionId   int    `json:"detectionId"`
	DetectionName string `json:"detectionName"`
	Level         int    `json:"level"`
	LevelName     string `json:"levelName"`
	Acronym       string `json:"acronym"`
}

func GetRoleExplore(gid int, sid, rid string, cid, code int, account config.Account) (*GetRoleExploreData, error) {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "serverId": sid, "roleId": rid, "channelId": strconv.Itoa(cid), "countryCode": strconv.Itoa(code)})
	return Exec[*GetRoleExploreData](req, "POST", "/gamer/roleBox/aki/exploreIndex")
}
