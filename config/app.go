// Package config 站点配置信息
package config

import pkConfig "go-web/pkg/config"

func init() {
	pkConfig.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 应用名称
			"name": pkConfig.Env("APP_NAME", "go-web"),

			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": pkConfig.Env("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": pkConfig.Env("APP_DEBUG", false),

			// 应用服务端口
			"port": pkConfig.Env("APP_PORT", "3000"),

			// 加密会话、JWT 加密
			"key": pkConfig.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": pkConfig.Env("APP_URL", "http://localhost:3000"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": pkConfig.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}