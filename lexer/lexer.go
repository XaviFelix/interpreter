package lexer

import (
	"main/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (point to current char)
	readPosition int  // curren reading position in input (after current char)
	ch           byte // current char under examination
}

// Return a pointer to the Lexer "instance" where only input was set to something
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // this will essentially initialize it

	return l
}

// All it does is read a char
func (l *Lexer) readChar() {
	// If read position is greater than the input string then ch is now 0
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ascii code for 'NUL'
	} else {
		l.ch = l.input[l.readPosition]
	}

	// move the pointers one
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
