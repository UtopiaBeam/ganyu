package lexer

import "fmt"

type Lexer interface {
	Tokenize(source string) ([]Token, error)
}

type lexer struct{}

var _ Lexer = (*lexer)(nil)

func NewLexer() Lexer {
	return &lexer{}
}

func (l *lexer) Tokenize(source string) ([]Token, error) {
	var (
		pos int

		row    = 1
		col    = 1
		tokens = make([]Token, 0)
	)

	for pos < len(source) {
		// if the current character is newline, skip and update position.
		if source[pos] == '\n' {
			pos++
			row, col = row+1, 1
			continue
		}

		// skip if it's skippable
		if step, skippable := matchSkippableToken(source[pos:]); skippable {
			pos += step
			col += step
			continue
		}

		matched := false
		for _, matcher := range matchers {
			if v := matcher.Matcher(source[pos:]); v != "" {
				matched = true
				pos += len(v)
				col += len(v)
				tokens = append(tokens, Token{Kind: matcher.Kind, Value: v})
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("invalid token at line %d col %d", row, col)
		}
	}

	return tokens, nil
}
