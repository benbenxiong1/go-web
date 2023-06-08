package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

// ValidateFunc 定义验证函数类型
type ValidateFunc func(interface{}) map[string][]string

// Validate 验证请求参数
func Validate(c *gin.Context, data interface{}, handle ValidateFunc) bool {
	// 解析json请求 支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(data); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
		fmt.Println("Validate ShouldBind err:", err.Error())
		return false
	}

	// 表单验证
	errs := handle(data)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"error":   errs,
		})
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, message govalidator.MapData) map[string][]string {

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
