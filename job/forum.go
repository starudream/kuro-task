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

type SignForumRecord struct {
	GameId    int
	GameName  string
	HasSigned bool
}

type SignForumRecords []SignForumRecord

func (rs SignForumRecords) Name() string {
	return "库街区每日任务"
}

func (rs SignForumRecords) Success() string {
	vs := []string{rs.Name() + "完成"}
	for i := 0; i < len(rs); i++ {
		vs = append(vs,
			fmt.Sprintf("在版区【%s】", rs[i].GameName),
			fmt.Sprintf(" 打卡成功"),
		)
	}
	return strings.Join(vs, "\n")
}

func SignForum(account config.Account) (SignForumRecords, error) {
	_, err := kuro.GetUser(account)
	if err != nil {
		return nil, fmt.Errorf("get user error: %w", err)
	}
	return SignForumGames(account)
}

func SignForumGames(account config.Account) (SignForumRecords, error) {
	var records []SignForumRecord
	for id, name := range kuro.GameNames {
		record, err := SignForumGame(kuro.GameIdMC, account)
		record.GameId = id
		record.GameName = name
		slog.Info("sign forum record: %+v", record)
		if err != nil {
			slog.Error("sign forum error: %w", err)
			continue
		}
		records = append(records, record)
	}
	slices.SortFunc(records, func(a, b SignForumRecord) int {
		return cmp.Compare(a.GameId, b.GameId)
	})
	return records, nil
}

func SignForumGame(gid int, account config.Account) (record SignForumRecord, err error) {
	today, err := kuro.GetSignForum(gid, account)
	if err != nil {
		err = fmt.Errorf("get sign forum error: %w", err)
		return
	}

	record.HasSigned = today.HasSignIn

	if today.HasSignIn {
		return
	}

	_, err = kuro.SignForum(gid, account)
	if err != nil {
		err = fmt.Errorf("sign forum error: %w", err)
		return
	}

	// todo: daily task for forum posts

	return
}
