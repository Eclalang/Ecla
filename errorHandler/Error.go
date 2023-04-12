package errorHandler

import "fmt"

// Error is the error struct of ecla errors.
type Error struct {
	Line  int
	Col   int
	Msg   string
	Level Level
}

// String returns the string representation of an error.
func (e Error) String() string {
	return fmt.Sprintf("%s Line: %d, Col: %d\n%s", e.Level, e.Line, e.Col, e.Msg)
}
