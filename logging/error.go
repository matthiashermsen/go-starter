package logging

func Error(message string, args ...any) {
	logger.Error(message, args...)
}
