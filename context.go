package muick

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/adibfahimi/muick/types"
)

type Map map[string]string

type Ctx struct {
	Request  types.HttpRequest
	Response types.HttpResponse
}

func (c *Ctx) SendString(body string) error {
	c.Response.Body = body

	return nil
}

func (c *Ctx) JSON(body interface{}) error {
	json, err := json.Marshal(body)
	if err != nil {
		return err
	}

	c.Response.Body = string(json)
	c.Response.Headers["Content-Type"] = "application/json"

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

func (c *Ctx) BodyParser(v interface{}) error {
	if c.Request.Body == "" {
		return errors.New("request body is empty")
	}

	fmt.Println(c.Request.Body)

	return json.Unmarshal([]byte(c.Request.Body), v)
}
