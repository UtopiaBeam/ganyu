package lexer

import "fmt"

type Error struct {
	Token byte
	Row   int
	Col   int
}

var _ error = (*Error)(nil)

func (e Error) Error() string {
	return fmt.Sprintf("invalid token %c at row %d, col %d", e.Token, e.Row, e.Col)
}
