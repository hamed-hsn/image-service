package protocol

type Logger interface {
	Info(string, ...any)
	Warning(string, ...any)
	Debug(string, ...any)
	Error(string, ...any)
}
