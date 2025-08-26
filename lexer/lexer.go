package lexer

type Lexer struct {
	input string
	pos int
	inBold   bool
    inItalic bool
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) nextChar() rune {
	if l.pos >= len(l.input) {
		return 0
	}

	ch := rune(l.input[l.pos])
	l.pos++
	return ch
}

func (l *Lexer) peekChar() rune {
	if l.pos >= len(l.input) {
		return 0
	}

	return rune(l.input[l.pos])
}

func (l *Lexer) NextToken() Token {
	ch := l.nextChar()
	
	switch ch {
		case '#':
			return Token{Type: TOKEN_HEADING, Value: "#"}
		case '*':
			if l.peekChar() == '*' {
				l.nextChar()
				if l.inBold {
					l.inBold = false
					return Token{Type: TOKEN_BOLD_END, Value: "**"}
				} else {
					l.inBold = true
					return Token{Type: TOKEN_BOLD_START, Value: "**"}
				}
			} else {
				if l.inItalic {
					l.inItalic = false
					return Token{Type: TOKEN_ITALIC_END, Value: "*"}
				} else {
					l.inItalic = true
					return Token{Type: TOKEN_ITALIC_START, Value: "*"}
				}
			}
		case '\n':
			return Token{Type: TOKEN_NEWLINE, Value: "\\n"}
		case 0:
			return Token{Type: TOKEN_EOF, Value: ""}
		default:
			start := l.pos - 1
			for {
				if l.peekChar() == '#' || l.peekChar() == '*' || l.peekChar() == '\n' || l.peekChar() == 0 {
					break
				}
				l.nextChar()
			}
			return Token{
				Type: TOKEN_TEXT,
				Value: l.input[start: l.pos],
			}
	}
}

func (l *Lexer) Lex() []Token {
	tokens := []Token{}
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == TOKEN_EOF {
			break
		}
	}
	return tokens
}