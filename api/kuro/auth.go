package kuro

import (
	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/gh"
)

type SendPhoneCodeData struct {
	GeeTest bool `json:"geeTest"`
}

type GeeTestData struct {
	CaptchaId     string `json:"captcha_id,omitempty"`
	LotNumber     string `json:"lot_number,omitempty"`
	PassToken     string `json:"pass_token,omitempty"`
	GenTime       string `json:"gen_time,omitempty"`
	CaptchaOutput string `json:"captcha_output,omitempty"`
}

func SendPhoneCodeGeeTest(phone string, geeTest *GeeTestData) (*SendPhoneCodeData, error) {
	data := "{}"
	if geeTest == nil {
		data = json.MustMarshalString(geeTest)
	}
	req := R().SetFormData(gh.MS{"mobile": phone, "geeTestData": data})
	return Exec[*SendPhoneCodeData](req, "POST", "/user/getSmsCode")
}

type LoginByPhoneCodeData struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

func LoginByPhoneCode(phone, code, devCode string) (*LoginByPhoneCodeData, error) {
	req := R().SetBody(gh.M{"mobile": phone, "code": code, "devCode": devCode})
	return Exec[*LoginByPhoneCodeData](req, "POST", "/user/sdkLogin")
}
