// Package requests  user : benbenxiong  time : 2023-07-2023/6/7 21:21:33
package requests

import (
	"github.com/thedevsaddam/govalidator"
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

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}

	return govalidator.New(opts).ValidateStruct()
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

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
