package errorHandler

import (
	"log"
	"time"
)

type ErrorHandler struct {
	ErrorChannel chan Error
	Errors       []Error
}

func NewHandler() *ErrorHandler {
	tempChan := make(chan Error)
	return &ErrorHandler{
		ErrorChannel: tempChan,
		Errors:       []Error{},
	}
}

func Send(Line, Col int, Message string, LogLevel Level, ErrorChan chan Error) {
	tempErr := Error{
		Line:  Line,
		Col:   Col,
		Msg:   Message,
		Level: LogLevel,
	}
	ErrorChan <- tempErr
}

func (ErrHandler *ErrorHandler) HandleError() {
	select {
	case err := <-ErrHandler.ErrorChannel:
		switch err.Level {
		case LevelWarning:
			log.Println("[" + time.Now().String() + "]" + LevelToString(err.Level) + " : " + err.Msg)
		case LevelError:
			log.Println("[" + time.Now().String() + "]" + LevelToString(err.Level) + " : " + err.Msg)
		case LevelFatal:
			panic("[" + time.Now().String() + "]" + LevelToString(err.Level) + " : " + err.Msg)

		}

	}
}
