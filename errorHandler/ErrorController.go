package errorHandler

import (
	"fmt"
	"log"
	"os"
)

// eclaExit is a variable to store the os.Exit function only modified for testing purposes.
var eclaExit = os.Exit
var oldExit = os.Exit

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
func (e *ErrorHandler) HandleError(Line, Col int, Message string, LogLevel Level) {
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

// HookExit is used for testing purpose it hooks the eclaExit variable to the function passed as parameter
func (e *ErrorHandler) HookExit(f func(int)) {
	oldExit = eclaExit
	eclaExit = f
}

// RestoreExit is used for testing purpose it restore the hook of eclaExit
func (e *ErrorHandler) RestoreExit() {
	eclaExit = oldExit
}

func panicEcla(err Error) {
	fmt.Println(err)
	eclaExit(1)
}
