package cofig

import "gohub/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 名称
			"name": config.Env("APP_NAME", "Gohub"),
			// 环境
			"env": config.Env("APP_ENV", "production"),
			// 调试
			"debug": config.Env("APP_DEBUG", false),
			// port
			"port": config.Env("APP_PORT", "3000"),
			// 加密会话，JWT
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			// url
			"url": config.Env("APP_URL", "http://localhost:3000"),
			// 时区
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
