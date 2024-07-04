package job

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/kuro-task/api/kuro"
	"github.com/starudream/kuro-task/config"
)

const (
	PostView  = 3
	PostLike  = 5
	PostShare = 1
	PostLoop  = 3
)

var ForumByGame = map[int]int{
	kuro.GameIdPNS: kuro.ForumIdPNS3,
	kuro.GameIdMC:  kuro.ForumIdMC10,
}

type SignForumRecord struct {
	GameId    int
	GameName  string
	HasSigned bool

	PostView  int
	PostLike  int
	PostShare int
	LoopCount int
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
		record, err := SignForumGame(id, account)
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
		goto post
	}

	_, err = kuro.SignForum(gid, account)
	if err != nil {
		err = fmt.Errorf("sign forum error: %w", err)
		return
	}

post:

	fid := ForumByGame[gid]

	record.LoopCount++

	posts, err := kuro.ListPost(gid, fid, record.LoopCount, account)
	if err != nil {
		err = fmt.Errorf("list post error: %w", err)
		return
	}

	for i := 0; i < len(posts.PostList); i++ {
		p := posts.PostList[i]
		if record.PostView < PostView {
			_, e := kuro.GetPost(p.PostId, account)
			if e != nil {
				slog.Error("get post error: %v", e)
				continue
			}
			record.PostView++
			time.Sleep(100 * time.Millisecond)
		}
		if record.PostLike < PostLike && p.IsLike == 0 {
			e := kuro.LikePost(gid, fid, p.PostId, p.UserId, false, account)
			if e != nil {
				slog.Error("like post error: %v", e)
				continue
			}
			time.Sleep(100 * time.Millisecond)
			e = kuro.LikePost(gid, fid, p.PostId, p.UserId, true, account)
			if e != nil {
				slog.Error("unlike post error: %v", e)
				continue
			}
			record.PostLike++
			time.Sleep(100 * time.Millisecond)
		}
		if record.PostShare < PostShare {
			e := kuro.SharePost(gid, account)
			if e != nil {
				slog.Error("share post error: %v", e)
				continue
			}
			record.PostShare++
			time.Sleep(100 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)
	}

	if record.LoopCount < PostLoop && (record.PostView < PostView || record.PostLike < PostLike || record.PostShare < PostShare) {
		goto post
	}

	return
}
