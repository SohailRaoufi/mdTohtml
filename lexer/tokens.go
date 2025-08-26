package lexer

type TokenType string

const (
    TOKEN_HEADING     TokenType = "HEADING"
    TOKEN_BOLD_START  TokenType = "BOLD_START"
    TOKEN_BOLD_END    TokenType = "BOLD_END"
    TOKEN_ITALIC_START TokenType = "ITALIC_START"
    TOKEN_ITALIC_END   TokenType = "ITALIC_END"
    TOKEN_TEXT        TokenType = "TEXT"
    TOKEN_NEWLINE     TokenType = "NEWLINE"
    TOKEN_EOF         TokenType = "EOF"
)

type Token struct {
	Type TokenType
	Value string
}