package muick

import (
	"encoding/json"

	"github.com/adibfahimi/muick/types"
)

type Map map[string]string

type Ctx struct {
	Request  types.HttpRequest
	Response types.HttpResponse
}

func (c *Ctx) SendString(body string) error {
	c.Response = types.HttpResponse{
		Body: body,
	}
	return nil
}

func (c *Ctx) JSON(body interface{}) error {
	json, err := json.Marshal(body)
	if err != nil {
		return err
	}

	c.Response = types.HttpResponse{
		Body: string(json),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return nil
}

func (c *Ctx) Append(header, value string) {
	c.Response.Headers[header] = value
}

func (c *Ctx) Accepts(accept string) {
	c.Response.Headers["Accept"] = accept
}

func (c *Ctx) Status(code int) *Ctx {
	c.Response.Status = code
	return c
}

func (c *Ctx) SendStatus(code int) {
	c.Response.Status = code
}
