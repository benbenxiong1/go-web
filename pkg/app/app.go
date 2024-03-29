package app

import (
	"go-web/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Env("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

func TImeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

func Url(path string) string {
	return config.Get("app.url") + path
}

func V1Url(path string) string {
	return Url("/v1/" + path)
}
