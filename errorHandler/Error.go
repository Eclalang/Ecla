package errorHandler

type Error struct {
	Line  int
	Col   int
	Msg   string
	Level Level
}
