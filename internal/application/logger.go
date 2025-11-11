package application

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, arg ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}
