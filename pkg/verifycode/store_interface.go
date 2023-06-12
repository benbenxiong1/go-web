package verifycode

type Store interface {
	// Set 设置验证码值
	Set(id string, value string) bool

	// Get 获取验证码值
	Get(id string, clear bool) string

	// Verify 检查验证码
	Verify(id, answer string, clear bool) bool
}
