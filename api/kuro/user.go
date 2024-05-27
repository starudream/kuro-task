package kuro

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/kuro-task/config"
)

type User struct {
	Mine Mine `json:"mine"`
}

type Mine struct {
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
	Mobile       string `json:"mobile"`
	Status       int    `json:"status"`
	Gender       int    `json:"gender"`
	GoldNum      int    `json:"goldNum"`
	IpRegion     string `json:"ipRegion"`
	RegisterTime string `json:"registerTime"`
	Signature    string `json:"signature"`
}

func GetUser(account config.Account) (*User, error) {
	req := R(account).SetFormData(gh.MS{"type": "1", "searchType": "2"})
	return Exec[*User](req, "POST", "/user/mine")
}
