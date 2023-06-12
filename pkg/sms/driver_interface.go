package sms

type Driver interface {
	Send(phone string, msg Message, config map[string]string) bool
}
