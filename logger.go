package gchat

type Logger interface {
	Info()
	Warn()
	Error()
	Debug()
}
