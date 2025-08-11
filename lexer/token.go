package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type Token struct {
	Kind  TokenKind
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%v(\"%v\")", t.Kind, t.Value)
}

type TokenKind string

const (
	// Keywords
	TokenKindType         TokenKind = "TYPE"
	TokenKindEnum         TokenKind = "ENUM"
	TokenKindInt          TokenKind = "INT"
	TokenKindFloat        TokenKind = "FLOAT"
	TokenKindBool         TokenKind = "BOOL"
	TokenKindString       TokenKind = "STRING"
	TokenKindExclaimation TokenKind = "EXCLAIMATION"

	TokenKindOpenCurly    TokenKind = "OPEN_CURLY"
	TokenKindCloseCurly   TokenKind = "CLOSE_CURLY"
	TokenKindOpenBracket  TokenKind = "OPEN_BRACKET"
	TokenKindCloseBracket TokenKind = "CLOSE_BRACKET"

	TokenKindComma TokenKind = "COMMA"
	TokenKindColon TokenKind = "COLON"

	TokenKindStringLiteral TokenKind = "STRING_LITERAL"
	TokenKindModifier      TokenKind = "MODIFIER"
)

type TokenMatcher struct {
	Kind TokenKind

	// Matcher returns the matched string (empty if none).
	Matcher func(string) string
}

var matchers = []TokenMatcher{
	{Kind: TokenKindType, Matcher: newKeywordMatcher("type")},
	{Kind: TokenKindEnum, Matcher: newKeywordMatcher("enum")},
	{Kind: TokenKindInt, Matcher: newKeywordMatcher("int")},
	{Kind: TokenKindFloat, Matcher: newKeywordMatcher("float")},
	{Kind: TokenKindBool, Matcher: newKeywordMatcher("bool")},
	{Kind: TokenKindString, Matcher: newKeywordMatcher("string")},
	{Kind: TokenKindExclaimation, Matcher: newKeywordMatcher("!")},
	{Kind: TokenKindOpenCurly, Matcher: newKeywordMatcher("{")},
	{Kind: TokenKindCloseCurly, Matcher: newKeywordMatcher("}")},
	{Kind: TokenKindOpenBracket, Matcher: newKeywordMatcher("[")},
	{Kind: TokenKindCloseBracket, Matcher: newKeywordMatcher("]")},
	{Kind: TokenKindComma, Matcher: newKeywordMatcher(",")},
	{Kind: TokenKindColon, Matcher: newKeywordMatcher(":")},
	{Kind: TokenKindStringLiteral, Matcher: matchStringLiteral},
	{Kind: TokenKindModifier, Matcher: matchModifier},
}

func matchStringLiteral(source string) string {
	// Attempt to split source by double-quote (").
	// The valid string literal should be splited into 3 parts.
	// For example, the string "text" should be splited into ["", "text", ""]
	strs := strings.SplitN(source, "\"", 3)
	if len(strs) != 3 || strs[0] != "" {
		return ""
	}
	return strs[1]
}

func newKeywordMatcher(tok string) func(string) string {
	return func(source string) string {
		if strings.HasPrefix(source, tok) {
			return tok
		}
		return ""
	}
}

func matchModifier(source string) string {
	// Allowed symbols are a-z, A-Z, - and _
	for i, c := range source {
		switch {
		case c == '-', c == '_',
			unicode.IsLetter(c):
			// continue
		default:
			return source[:i]
		}
	}
	return source
}

func matchSkippableToken(source string) (int, bool) {
	if source == "" {
		return 0, false
	}
	switch source[0] {
	case ' ':
		return 1, true
	case '\t':
		return 4, true
	default:
		return 0, false
	}
}
