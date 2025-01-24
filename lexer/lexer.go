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

	l.skipWhitespace()

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// all white spaces are ignored
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// if the token is an identifer then this function
// will read each char from the input and extract the
// identifer and create a token out of it
func (l *Lexer) readIdentifier() string {

	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // this is a slice: input[start:end] start is inclusive, end is exclusive
}

// Checks to see if the current byte is a letter represented in the ascii table
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// creates a new token "instance"
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
