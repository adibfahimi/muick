package parser

import (
	"strings"

	"github.com/adibfahimi/muick/types"
)

func ParseRequest(input string) types.HttpRequest {
	l := NewLexer(input)
	req := types.HttpRequest{}
	req.Method = strings.TrimSpace(l.readUntil(' '))
	req.Path = strings.TrimSpace(l.readUntilSequence(" HTTP"))
	req.Version = l.readLine()
	req.Headers = l.readHeaders()
	req.Body = l.readBody()
	return req
}
