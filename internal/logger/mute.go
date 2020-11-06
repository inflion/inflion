package logger

// MuteLogger returns a `Logger` does noting.
var MuteLogger Logger

func init() {
	MuteLogger = &muteLogger{}
}

type muteLogger struct{}

func (m *muteLogger) Debug(args ...interface{}) {
	// Do nothing
}

func (m *muteLogger) Info(args ...interface{}) {
	// Do nothing
}

func (m *muteLogger) Error(args ...interface{}) {
	// Do nothing
}

func (m *muteLogger) Fatal(args ...interface{}) {
	// Do nothing
}

func (m *muteLogger) DebugWith(msg string, keysAndValues ...interface{}) {
	// Do nothing
}

func (m *muteLogger) InfoWith(msg string, keysAndValues ...interface{}) {
	// Do nothing
}

func (m *muteLogger) ErrorWith(msg string, keysAndValues ...interface{}) {
	// Do nothing
}

func (m *muteLogger) FatalWith(msg string, keysAndValues ...interface{}) {
	// Do nothing
}
