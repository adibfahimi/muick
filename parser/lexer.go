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
	start := l.position
	for l.ch != until && l.ch != 0 {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readLine() string {
	return l.readUntil('\n')
}

func (l *Lexer) readHeaders() map[string]string {
	headers := make(map[string]string)
	for l.ch != '\n' && l.ch != 0 {
		key := strings.TrimSpace(l.readUntil(':'))
		l.readChar()
		value := strings.TrimSpace(l.readLine())
		headers[key] = value
	}
	return headers
}

func (l *Lexer) readBody() string {
	l.readUntilSequence("\r\n\r\n")

	return l.input[l.position+4:]
}

func (l *Lexer) readUntilSequence(sequence string) string {
	start := l.position
	for {
		if strings.HasPrefix(l.input[l.position:], sequence) {
			break
		}
		if l.readPosition >= len(l.input) {
			l.ch = 0
			break
		}
		l.readChar()
	}
	return l.input[start:l.position]
}
