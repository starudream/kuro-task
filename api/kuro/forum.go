package kuro

import (
	"strconv"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/kuro-task/config"
)

type ListPostData struct {
	HasNext  int         `json:"hasNext"`
	PostList []*PostInfo `json:"postList"`
}

type PostInfo struct {
	GameId          int       `json:"gameId"`
	GameName        string    `json:"gameName"`
	GameForumId     int       `json:"gameForumId"`
	PostId          string    `json:"postId"`
	PostType        int       `json:"postType"`
	PostTitle       string    `json:"postTitle"`
	PostContent     any       `json:"postContent"`
	UserId          string    `json:"userId"`
	UserName        string    `json:"userName"`
	BrowseCount     string    `json:"browseCount"`
	CommentCount    int       `json:"commentCount"`
	LikeCount       int       `json:"likeCount"`
	IsFollow        int       `json:"isFollow"`
	IsLike          int       `json:"isLike"`
	IsLock          int       `json:"isLock"`
	IsPublisher     int       `json:"isPublisher"`
	CreateTimestamp Timestamp `json:"createTimestamp"`
}

func ListPost(gid, fid, page int, account config.Account) (*ListPostData, error) {
	// searchType 1最新发布 2最新回复 3默认
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "forumId": strconv.Itoa(fid), "searchType": "1", "pageIndex": strconv.Itoa(page), "pageSize": "20"})
	return Exec[*ListPostData](req, "POST", "/forum/list")
}

func GetPost(id string, account config.Account) (*GetPostData, error) {
	req := R(account).SetFormData(gh.MS{"postId": id})
	return Exec[*GetPostData](req, "POST", "/forum/getPostDetail")
}

type GetPostData struct {
	PostDetail *PostInfo `json:"postDetail"`
	IsCollect  int       `json:"isCollect"`
	IsFollow   int       `json:"isFollow"`
	IsLike     int       `json:"isLike"`
}

func LikePost(gid, fid int, pid, uid string, cancel bool, account config.Account) error {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid), "forumId": strconv.Itoa(fid), "likeType": "1", "operateType": gh.Ternary(cancel, "2", "1"), "postId": pid, "postType": "1", "toUserId": uid})
	_, err := Exec[bool](req, "POST", "/forum/like")
	return err
}

func SharePost(gid int, account config.Account) error {
	req := R(account).SetFormData(gh.MS{"gameId": strconv.Itoa(gid)})
	_, err := Exec[*GetPostData](req, "POST", "/encourage/level/shareTask")
	return err
}
