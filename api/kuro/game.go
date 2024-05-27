package kuro

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/kuro-task/config"
)

type Game3Widget struct {
	GameId         int                `json:"gameId"`
	UserId         int                `json:"userId"`
	ServerTime     int                `json:"serverTime"`
	ServerId       string             `json:"serverId"`
	ServerName     string             `json:"serverName"`
	SignInTxt      string             `json:"signInTxt"`
	HasSignIn      bool               `json:"hasSignIn"`
	RoleId         string             `json:"roleId"`
	RoleName       string             `json:"roleName"`
	EnergyData     *Game3WidgetItem   `json:"energyData"`
	LivenessData   *Game3WidgetItem   `json:"livenessData"`
	BattlePassData []*Game3WidgetItem `json:"battlePassData"`
}

type Game3WidgetItem struct {
	Name             string `json:"name"`
	RefreshTimeStamp int    `json:"refreshTimeStamp"`
	ExpireTimeStamp  int    `json:"expireTimeStamp"`
	Status           int    `json:"status"`
	Cur              int    `json:"cur"`
	Total            int    `json:"total"`
}

func GetGame3Widget(account config.Account) (*Game3Widget, error) {
	req := R(account).SetFormData(gh.MS{"type": "1", "sizeType": "2"})
	return Exec[*Game3Widget](req, "POST", "/gamer/widget/game3/getData")
}

type Role struct {
	UserId     int    `json:"userId"`
	GameId     int    `json:"gameId"`
	ServerId   string `json:"serverId"`
	ServerName string `json:"serverName"`
	RoleId     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	IsDefault  bool   `json:"isDefault"`
}

func ListRole(gid string, account config.Account) ([]*Role, error) {
	req := R(account).SetFormData(gh.MS{"gameId": gid})
	return Exec[[]*Role](req, "POST", "/gamer/role/list")
}
