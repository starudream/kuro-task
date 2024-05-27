package kuro

import (
	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/gh"
)

type SendPhoneCodeData struct {
	GeeTest bool `json:"geeTest"`
}

func SendPhoneCode(phone string) (*SendPhoneCodeData, error) {
	req := R().SetFormData(gh.MS{"mobile": phone, "geeTestData": "{}"})
	return Exec[*SendPhoneCodeData](req, "POST", "/user/getSmsCode")
}

type GeeTestData struct {
	CaptchaId     string `json:"captcha_id"`
	LotNumber     string `json:"lot_number"`
	PassToken     string `json:"pass_token"`
	GenTime       string `json:"gen_time"`
	CaptchaOutput string `json:"captcha_output"`
}

func SendPhoneCodeGeeTest(phone string, geeTest *GeeTestData) (*SendPhoneCodeData, error) {
	req := R().SetFormData(gh.MS{"mobile": phone, "geeTestData": json.MustMarshalString(geeTest)})
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
