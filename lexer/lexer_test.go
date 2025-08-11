package lexer

import (
	_ "embed"
	"reflect"
	"testing"
)

var (
	//go:embed 01-invalid-token.gan
	invalidSyntaxSource string

	//go:embed 02-valid-source.gan
	validSource string
)

func Test_LexerTokenize(t *testing.T) {
	type testCase struct {
		name    string
		source  string
		want    []Token
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "return error if token is invalid",
			source:  invalidSyntaxSource,
			want:    nil,
			wantErr: true,
		},
		{
			name:   "return tokens",
			source: validSource,
			want: []Token{
				{Kind: TokenKindType, Value: "type"},
				{Kind: TokenKindModifier, Value: "APIResponse"},
				{Kind: TokenKindOpenCurly, Value: "{"},
				{Kind: TokenKindModifier, Value: "code"},
				{Kind: TokenKindColon, Value: ":"},
				{Kind: TokenKindString, Value: "string"},
				{Kind: TokenKindExclaimation, Value: "!"},
				{Kind: TokenKindModifier, Value: "message"},
				{Kind: TokenKindColon, Value: ":"},
				{Kind: TokenKindString, Value: "string"},
				{Kind: TokenKindCloseCurly, Value: "}"},
			},
			wantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLexer().Tokenize(tt.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr %v, got %v", tt.wantErr, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}
