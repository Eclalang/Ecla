package errorHandler

import "fmt"

type Error struct {
	Line  int
	Col   int
	Msg   string
	Level Level
}

func (e Error) String() string {
	return fmt.Sprintf("%s Line: %d, Col: %d\n%s", e.Level, e.Line, e.Col, e.Msg)
}
