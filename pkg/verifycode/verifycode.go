package verifycode

import (
	"fmt"
	"go-web/pkg/app"
	"go-web/pkg/config"
	mail "go-web/pkg/email"
	"go-web/pkg/helpers"
	"go-web/pkg/logger"
	"go-web/pkg/redis"
	"go-web/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once

var internalVerifyCode *VerifyCode

func NewVerify() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

// SendSMS 发送短信验证码，调用示例：
func (v *VerifyCode) SendSMS(phone string) bool {
	// 生成验证码
	code := v.generateVerifyCode(phone)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}

	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

func (v *VerifyCode) SendEmail(email string) bool {
	// 生成验证码
	code := v.generateVerifyCode(email)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasSuffix(email, config.GetString("verifycode.debug_email_suffix")) {
		return true
	}

	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)
	// 发送邮件
	return mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email 验证码",
		Html:    []byte(content),
	})
}

// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (v *VerifyCode) CheckAnswer(id string, value string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{id: value})

	// 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	if !app.IsProduction() &&
		(strings.HasSuffix(id, config.GetString("verifycode.debug_email_suffix")) ||
			strings.HasPrefix(id, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}

	return v.Store.Verify(id, value, false)
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (v *VerifyCode) generateVerifyCode(key string) string {
	// 生成随机码
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	// 为方便开发，本地环境使用固定验证码
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	v.Store.Set(key, code)
	return code
}
