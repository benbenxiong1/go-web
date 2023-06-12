package sms

import (
	"go-web/pkg/config"
	"sync"
)

// Message 是短信的结构体
type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

// SMS 是我们发送短信的操作类
type SMS struct {
	Driver Driver
}

// once 单利模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 实例化SMS
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}

func (s *SMS) Send(phone string, message Message) bool {
	return s.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
