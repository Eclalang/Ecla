package errorHandler

type Level int

const (
	// LevelWarning is the warning level
	LevelWarning Level = iota
	// LevelError is the error level
	LevelError
	// LevelFatal is the fatal level
	LevelFatal
)

func LevelToString(level Level) string {
	switch level {
	case LevelWarning:
		return "Warning"
	case LevelError:
		return "Error"
	case LevelFatal:
		return "Fatal"
	default:
		return ""
	}
}

func StringToLevel(level string) Level {
	switch level {
	case "Warning":
		return LevelWarning
	case "Error":
		return LevelError
	case "Fatal":
		return LevelFatal
	default:
		return -1
	}
}

func (l Level) String() string {
	return LevelToString(l)
}
