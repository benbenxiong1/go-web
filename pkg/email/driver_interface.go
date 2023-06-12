package email

type Driver interface {
	Send(email Email, config map[string]string) bool
}
