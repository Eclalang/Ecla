package errorHandler

import (
	"testing"
)

func TestError(t *testing.T) {
	e := Error{
		Line:  1,
		Col:   2,
		Msg:   "Test",
		Level: LevelError,
	}
	if e.String() != "Error Line: 1, Col: 2\nTest" {
		t.Errorf("Error.String() returned wrong string")
	}
}

func TestLevelToString(t *testing.T) {
	if LevelToString(LevelWarning) != "Warning" {
		t.Errorf("LevelToString(LevelWarning) returned wrong string")
	}
	if LevelToString(LevelError) != "Error" {
		t.Errorf("LevelToString(LevelError) returned wrong string")
	}
	if LevelToString(LevelFatal) != "Fatal Error" {
		t.Errorf("LevelToString(LevelFatal) returned wrong string")
	}
	if LevelToString(-1) != "" {
		t.Errorf("LevelToString(-1) returned wrong string")
	}
}

func TestStringToLevel(t *testing.T) {
	if StringToLevel("Warning") != LevelWarning {
		t.Errorf("StringToLevel(\"Warning\") returned wrong level")
	}
	if StringToLevel("Error") != LevelError {
		t.Errorf("StringToLevel(\"Error\") returned wrong level")
	}
	if StringToLevel("Fatal") != LevelFatal {
		t.Errorf("StringToLevel(\"Fatal\") returned wrong level")
	}
	if StringToLevel("Test") != -1 {
		t.Errorf("StringToLevel(\"Test\") returned wrong level")
	}
}

func TestNewHandler(t *testing.T) {
	e := NewHandler()
	if len(e.Errors) != 0 {
		t.Errorf("NewHandler() returned wrong ErrorHandler")
	}
}

func TestHandleError(t *testing.T) {
	e := NewHandler()
	e.HandleError(1, 2, "Test", LevelError)
	if len(e.Errors) != 1 {
		t.Errorf("HandleError() did not append the error")
	}
	if e.Errors[0].Line != 1 || e.Errors[0].Col != 2 || e.Errors[0].Msg != "Test" || e.Errors[0].Level != LevelError {
		t.Errorf("HandleError() appended the wrong error")
	}
	e.HandleError(3, 4, "Test2", LevelWarning)
	if len(e.Errors) != 2 {
		t.Errorf("HandleError() did not append the error")
	}
	if e.Errors[1].Line != 3 || e.Errors[1].Col != 4 || e.Errors[1].Msg != "Test2" || e.Errors[1].Level != LevelWarning {
		t.Errorf("HandleError() appended the wrong error")
	}
	var ok bool
	eclaExit = func(code int) {
		ok = code == 1
	}
	e.HandleError(5, 6, "Test3", LevelFatal)
	if !ok {
		t.Errorf("HandleError() did not panic")
	}
	if len(e.Errors) != 3 {
		t.Errorf("HandleError() did not append the error")
	}
	if e.Errors[2].Line != 5 || e.Errors[2].Col != 6 || e.Errors[2].Msg != "Test3" || e.Errors[2].Level != LevelFatal {
		t.Errorf("HandleError() appended the wrong error")
	}
}

func TestPanicEcla(t *testing.T) {
	var ok bool
	eclaExit = func(code int) {
		ok = code == 1
	}
	panicEcla(Error{
		Line:  1,
		Col:   2,
		Msg:   "Test",
		Level: LevelFatal,
	})
	if !ok {
		t.Errorf("panicEcla() did not panic")
	}
}
