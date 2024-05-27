package config

import (
	"strings"

	"github.com/kr/pretty"

	"github.com/starudream/go-lib/core/v2/slog"
)

type Account struct {
	Phone   string `json:"phone"    yaml:"phone"`
	DevCode string `json:"dev_code" yaml:"dev_code"`
	Source  string `json:"source"   yaml:"source"`
	Version string `json:"version"  yaml:"version"`
	Token   string `json:"token"    yaml:"token"   table:",ignore"`
}

func AddAccount(account Account) {
	_cMu.Lock()
	defer _cMu.Unlock()
	u := false
	for i := range _c.Accounts {
		if _c.Accounts[i].Phone == account.Phone {
			_c.Accounts[i], u = account, true
		}
	}
	if !u {
		_c.Accounts = append(_c.Accounts, account)
	}
}

func UpdateAccount(phone string, cb func(account Account) Account) {
	_cMu.Lock()
	defer _cMu.Unlock()
	for i := range _c.Accounts {
		if _c.Accounts[i].Phone == phone {
			c := _c.Accounts[i]
			nc := cb(c)
			slog.Info("update account %s, diff: %s", phone, strings.Join(pretty.Diff(c, nc), ", "))
			_c.Accounts[i] = nc
			return
		}
	}
}

func GetAccount(phone string) (Account, bool) {
	accounts := C().Accounts
	for i := range accounts {
		if accounts[i].Phone == phone {
			return accounts[i], true
		}
	}
	return Account{}, false
}
