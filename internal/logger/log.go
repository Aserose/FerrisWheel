package logger

import "github.com/sirupsen/logrus"

type logger struct {
	log *logrus.Logger
}

func NewLogger() *logger {
	return &logger{
		log: logrus.New(),
	}
}

func (l *logger) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args)
}

func (l *logger) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l *logger) Print(args ...interface{}) {
	l.log.Print(args)
}

func (l *logger) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args)
}
