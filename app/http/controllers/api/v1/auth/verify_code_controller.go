package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/app/requests"
	"go-web/pkg/captcha"
	"go-web/pkg/logger"
	"go-web/pkg/response"
	"go-web/pkg/verifycode"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseApiController
}

func (v *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成图片验证码
	ids, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	response.Json(c, gin.H{
		"captcha_id":    ids,
		"captcha_image": b64s,
	})
}

func (v *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, request, requests.VerifyCodePhone); !ok {
		return
	}

	// 发送短信
	if ok := verifycode.NewVerify().SendSMS(request.Phone); !ok {
		response.Abort500(c, "短信发送失败。。。")
	} else {
		response.Success(c)
	}
}
