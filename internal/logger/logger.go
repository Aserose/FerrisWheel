package logger

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}
