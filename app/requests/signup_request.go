// Package requests  user : benbenxiong  time : 2023-07-2023/6/7 21:21:33
package requests

import (
	"github.com/thedevsaddam/govalidator"
	"go-web/app/requests/validators"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// ValidateSignupPhoneExist 验证手机号
func ValidateSignupPhoneExist(data interface{}) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义错误提示
	message := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return validate(data, rules, message)
}

func ValidateSignupEmailExist(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{
			"required", "min:4", "max:30", "email",
		},
	}

	message := govalidator.MapData{
		"email": []string{
			"required:email不能为空",
			"min:email最小长度不能小于4",
			"max:email最大长度不能大于30",
			"email:email格式错误",
		},
	}

	return validate(data, rules, message)
}

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	Name            string `json:"name,omitempty" valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
}

func SignupUsingPhone(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"phone":            []string{"require", "digits:11", "not_exists:user,phone"},
		"name":             []string{"require", "alpha_num", "between:3:20", "not_exists:user,name"},
		"password":         []string{"require", "min:6"},
		"password_confirm": []string{"require"},
		"verify_code":      []string{"require", "digits:6"},
	}

	message := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, message)

	_data := data.(SignupUsingPhoneRequest)
	// 验证2次密码是否一致
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)

	// 验证验证码是否正确
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
