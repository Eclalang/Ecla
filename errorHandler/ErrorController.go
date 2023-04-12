package errorHandler

import (
	"fmt"
	"log"
	"os"
)

// ErrorHandler is the error handler of ecla.
type ErrorHandler struct {
	Errors []Error
}

// NewHandler returns a new ErrorHandler.
func NewHandler() *ErrorHandler {
	return &ErrorHandler{
		Errors: []Error{},
	}
}

// HandleError handles an error.
func (e ErrorHandler) HandleError(Line, Col int, Message string, LogLevel Level) {
	err := Error{
		Line:  Line,
		Col:   Col,
		Msg:   Message,
		Level: LogLevel,
	}
	e.Errors = append(e.Errors, err)
	switch LogLevel {
	case LevelWarning:
		log.Println(LevelToString(LogLevel) + " : " + Message)
	case LevelError:
		log.Println(LevelToString(LogLevel) + " : " + Message)
	case LevelFatal:
		panicEcla(err)
	}
}

func panicEcla(err Error) {
	fmt.Println(err)
	os.Exit(1)
}
