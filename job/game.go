package job

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/kuro-task/api/kuro"
	"github.com/starudream/kuro-task/config"
)

type SignGameRecord struct {
	GameId     int
	GameName   string
	ServerName string
	RoleId     string
	RoleName   string
	HasSigned  bool
	Award      string
}

type SignGameRecords []SignGameRecord

func (rs SignGameRecords) Name() string {
	return "库街区游戏签到"
}

func (rs SignGameRecords) Success() string {
	vs := []string{rs.Name() + "完成"}
	for i := 0; i < len(rs); i++ {
		vs = append(vs, fmt.Sprintf("在游戏【%s】角色【%s】区服【%s】获得 %s", rs[i].GameName, rs[i].RoleName, rs[i].ServerName, rs[i].Award))
	}
	return strings.Join(vs, "\n")
}

func SignGame(account config.Account) (SignGameRecords, error) {
	_, err := kuro.GetUser(account)
	if err != nil {
		return nil, fmt.Errorf("get user error: %w", err)
	}
	roles1, err := kuro.ListRole(kuro.GameIdMC, account)
	if err != nil {
		return nil, fmt.Errorf("list role1 error: %w", err)
	}
	roles2, err := kuro.ListRole(kuro.GameIdPNS, account)
	if err != nil {
		return nil, fmt.Errorf("list role2 error: %w", err)
	}
	return SignGameRoles(append(roles1, roles2...), account)
}

func SignGameRoles(roles []*kuro.Role, account config.Account) (SignGameRecords, error) {
	var records []SignGameRecord
	for _, role := range roles {
		record, err := SignGameRole(role, account)
		slog.Info("sign game record: %+v", record)
		if err != nil {
			slog.Error("sign game error: %w", err)
			continue
		}
		records = append(records, record)
	}
	slices.SortFunc(records, func(a, b SignGameRecord) int {
		if a.GameId == b.GameId {
			return cmp.Compare(a.RoleId, b.RoleId)
		}
		return cmp.Compare(a.GameId, b.GameId)
	})
	return records, nil
}

func SignGameRole(role *kuro.Role, account config.Account) (record SignGameRecord, err error) {
	record.GameId = role.GameId
	record.ServerName = role.ServerName
	record.RoleId = role.RoleId
	record.RoleName = role.RoleName

	gameName, ok := kuro.GameNames[record.GameId]
	if !ok {
		err = fmt.Errorf("game id %d not support", role.GameId)
		return
	}

	record.GameName = gameName

	records, err := kuro.ListSignGameRecord(role.GameId, role.ServerId, role.RoleId, role.UserId, account)
	if err != nil {
		err = fmt.Errorf("list sign game record error: %w", err)
		return
	}

	today := records.Today()
	if len(today) > 0 {
		record.HasSigned = true
		record.Award = today.ShortString()
		return
	}

	signGameData, err := kuro.SignGame(role.GameId, role.ServerId, role.RoleId, role.UserId, account)
	if err != nil {
		if kuro.IsCode(err, kuro.CodeHasSigned) {
			record.HasSigned = true
		} else {
			err = fmt.Errorf("sign game error: %w", err)
			return
		}
	} else {
		list, err2 := kuro.ListSignGame(role.GameId, role.ServerId, role.RoleId, role.UserId, account)
		if err2 != nil {
			err2 = fmt.Errorf("list sign game error: %w", err2)
			return
		}
		record.Award = signGameData.TodayList.ShortStringByMap(list.GoodsMap())
	}

	return
}
