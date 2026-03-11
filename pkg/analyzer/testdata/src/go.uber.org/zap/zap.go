package zap

type Logger struct{}
type SugaredLogger struct{}

func NewProduction() (*Logger, error)   { return &Logger{}, nil }
func (l *Logger) Sugar() *SugaredLogger { return &SugaredLogger{} }

func (l *Logger) Info(msg string, fields ...interface{})  {}
func (l *Logger) Error(msg string, fields ...interface{}) {}
func (l *Logger) Warn(msg string, fields ...interface{})  {}
func (l *Logger) Debug(msg string, fields ...interface{}) {}

func (s *SugaredLogger) Info(msg string, args ...interface{})            {}
func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{})  {}
func (s *SugaredLogger) Infof(template string, args ...interface{})      {}
func (s *SugaredLogger) Error(msg string, args ...interface{})           {}
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {}
func (s *SugaredLogger) Errorf(template string, args ...interface{})     {}
func (s *SugaredLogger) Warn(msg string, args ...interface{})            {}
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{})  {}
func (s *SugaredLogger) Warnf(template string, args ...interface{})      {}
func (s *SugaredLogger) Debug(msg string, args ...interface{})           {}
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {}
func (s *SugaredLogger) Debugf(template string, args ...interface{})     {}
