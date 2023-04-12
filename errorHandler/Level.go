package errorHandler

// Level is the level of an error.
type Level int

const (
	// LevelWarning is the warning level
	LevelWarning Level = iota
	// LevelError is the error level
	LevelError
	// LevelFatal is the fatal level
	LevelFatal
)

// LevelToString returns the string representation of a level.
func LevelToString(level Level) string {
	switch level {
	case LevelWarning:
		return "Warning"
	case LevelError:
		return "Error"
	case LevelFatal:
		return "Fatal Error"
	default:
		return ""
	}
}

// StringToLevel returns the level of a string.
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

// String returns the string representation of a level.
func (l Level) String() string {
	return LevelToString(l)
}
