package kuro

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/kuro-task/config"
)

const (
	Addr          = "https://api.kurobbs.com"
	UserAgent     = "okhttp/3.11.0"
	AndroidName   = "com.kurogame.kjq"
	Version       = "2.2.0"
	SourceAndroid = "android"
	SourceH5      = "h5"

	CodeHasSigned = Code(1511)
)

const (
	GameIdPNS = 2 // 战双
	GameIdMC  = 3 // 鸣潮
)

const (
	ForumIdPNS2 = 2 // 推荐
	ForumIdPNS3 = 3 // 伊甸闲庭
	ForumIdPNS4 = 4 // 攻略
	ForumIdPNS5 = 5 // 同人

	ForumIdMC9  = 9  // 推荐
	ForumIdMC10 = 10 // 今州茶馆
	ForumIdMC11 = 11 // 同人
	ForumIdMC12 = 12 // 攻略
	ForumIdMC15 = 15 // 新手
)

const (
	GameMCCountryCodeHL = 1 // 瑝珑
)

var GameNames = map[int]string{
	GameIdPNS: "战双",
	GameIdMC:  "鸣潮",
}

type BaseResp[T any] struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`

	Data T `json:"data,omitempty"`
}

func (t *BaseResp[T]) GetCode() Code {
	if t == nil {
		return 0
	}
	return t.Code
}

func (t *BaseResp[T]) IsSuccess() bool {
	return t != nil && t.Code == 200
}

func (t *BaseResp[T]) String() string {
	return fmt.Sprintf("code: %d, message: %s", t.Code, t.Msg)
}

func IsCode(err error, code Code) bool {
	if err == nil {
		return false
	}
	e, ok1 := resty.AsRespErr(err)
	if ok1 {
		t, ok2 := e.Response.Result().(interface{ GetCode() Code })
		return ok2 && t.GetCode() == code
	}
	return false
}

func R(vs ...any) *resty.Request {
	r := resty.R().
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("User-Agent", UserAgent).
		SetHeader("X-Requested-With", AndroidName)
	// anyway, always set the headers
	m := gh.MS{
		"Source":  SourceH5,
		"Version": Version,
		"DevCode": "",
		"Token":   "",
	}
	for i := 0; i < len(vs); i++ {
		switch v := vs[i].(type) {
		case config.Account:
			if v.Source != "" {
				m["Source"] = v.Source
			}
			if v.Version != "" {
				m["Version"] = v.Version
			}
			m["DevCode"] = v.DevCode
			m["Token"] = v.Token
		}
	}
	r.SetHeaders(m)
	return r
}

func Exec[T any](r *resty.Request, method, path string) (t T, _ error) {
	res, err := resty.ParseResp[*BaseResp[any], *BaseResp[T]](
		r.SetError(&BaseResp[any]{}).SetResult(&BaseResp[T]{}).Execute(method, Addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[kuro] %w", err)
	}
	return res.Data, nil
}
