package parser

import "strings"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readUntil(until byte) string {
	position := l.position
	for l.ch != until && l.ch != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readHeaders() map[string]string {
	headers := map[string]string{}
	for l.ch != '\n' && l.ch != 0 {
		key := l.readUntil(':')
		l.readChar()
		value := l.readUntil('\n')
		headers[key] = value
		if l.ch == '\r' {
			l.readChar()
		}
		l.readChar()
	}
	return headers
}

func (l *Lexer) readUntilSequence(seq string) string {
	start := l.position
	for !strings.HasPrefix(l.input[l.position:], seq) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readLine() string {
	start := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[start:l.position]
}
