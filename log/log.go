package log

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type EmptyLogger struct{}

func (l *EmptyLogger) Debug(msg string, args ...any) {}
func (l *EmptyLogger) Info(msg string, args ...any)  {}
func (l *EmptyLogger) Warn(msg string, args ...any)  {}
func (l *EmptyLogger) Error(msg string, args ...any) {}
