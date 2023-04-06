package gchat

type Logger interface {
	Info()
	Warn()
	Error()
}

type logger struct {
}

var defaultLogger logger

func (l logger) Info() {
	//TODO implement me
	panic("implement me")
}

func (l logger) Warn() {
	//TODO implement me
	panic("implement me")
}

func (l logger) Error() {
	//TODO implement me
	panic("implement me")
}
