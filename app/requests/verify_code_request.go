package requests

import (
	"github.com/thedevsaddam/govalidator"
	"go-web/app/requests/validators"
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

	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaValue, errs)

	return errs
}

type VerifyCodeEmailRequest struct {
	CaptchaID    string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaValue string `json:"captcha_value,omitempty" valid:"captcha_value"`

	Email string `json:"email,omitempty" valid:"email"`
}

func VerifyCodeEmail(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"captcha_id":    []string{"required"},
		"captcha_value": []string{"required", "digits:6"},
		"email":         []string{"required", "min:4", "max:30", "email"},
	}

	message := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_value": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rules, message)

	_data := data.(*VerifyCodeEmailRequest)

	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaValue, errs)

	return errs
}
