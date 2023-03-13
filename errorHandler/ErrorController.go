package errorHandler

import (
	"fmt"
	"log"
	"os"
)

type ErrorHandler struct {
	Errors []Error
}

func NewHandler() *ErrorHandler {
	return &ErrorHandler{
		Errors: []Error{},
	}
}

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
