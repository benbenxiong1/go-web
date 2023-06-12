package requests

import (
	"github.com/thedevsaddam/govalidator"
	"go-web/pkg/captcha"
)

type VerifyCodePhoneRequest struct {
	CaptchaID    string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaValue string `json:"captcha_value,omitempty" valid:"captcha_value"`
	Phone        string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone 验证表单，返回长度等于零即通过
func VerifyCodePhone(data interface{}) map[string][]string {
	// 定义规则
	relus := govalidator.MapData{
		"captcha_id":    []string{"required"},
		"captcha_value": []string{"required", "digits:6"},
		"phone":         []string{"required", "digits:11"},
	}

	// 定义错误信息
	message := govalidator.MapData{
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_value": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	errs := validate(data, relus, message)

	//
	_data := data.(*VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaValue); !ok {
		errs["captcha_value"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}
