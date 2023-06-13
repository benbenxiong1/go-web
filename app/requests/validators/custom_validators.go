package validators

import (
	"go-web/pkg/captcha"
	"go-web/pkg/verifycode"
)

// ValidateCaptcha 自定义规则，验证『图片验证码』
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

// ValidatePasswordConfirm 自定义验证2次密码是否一致
func ValidatePasswordConfirm(password, passwordConfirm string, errors map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errors["password_confirm"] = append(errors["password_confirm"], "两次密码不一致")
	}
	return errors
}

func ValidateVerifyCode(phone string, verifyCode string, errors map[string][]string) map[string][]string {
	if ok := verifycode.NewVerify().CheckAnswer(phone, verifyCode); !ok {
		errors["verify_code"] = append(errors["verify_code"], "验证码错误")
	}
	return errors
}
