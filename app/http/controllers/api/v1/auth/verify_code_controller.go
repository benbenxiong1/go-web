package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/pkg/captcha"
	"go-web/pkg/logger"
	"net/http"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseApiController
}

func (v *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成图片验证码
	ids, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    ids,
		"captcha_image": b64s,
	})
}
