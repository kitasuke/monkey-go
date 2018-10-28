package lexer

import (
	"testing"

	"github.com/kitasuke/monkey-go/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}
		10 == 10;
		10 != 9;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LeftParen, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "result"},
		{token.Assign, "="},
		{token.Identifier, "add"},
		{token.LeftParen, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RightParen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Int, "5"},
		{token.LessThan, "<"},
		{token.Int, "10"},
		{token.GreaterThan, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.LeftParen, "("},
		{token.Int, "5"},
		{token.LessThan, "<"},
		{token.Int, "10"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Else, "else"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Int, "10"},
		{token.Equal, "=="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.NotEqual, "!="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
