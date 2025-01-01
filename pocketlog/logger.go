package pocketlog
// Logger is used to log information
type Logger struct {
	threshold Level
}
// New returns a logger, ready to log at the required threshold
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
// Debug formats and prints a message if the log leve is debug
func (l *Logger) Debugf(format string, args ...any) {}
// Infof formats and prints a message if the leg level is info
func (l *Logger) Infof(format string, args ...any) {}
// Errorf formats and prints a message if the log level is error
func (l *Logger) Errorf(format string, args ...any) {}