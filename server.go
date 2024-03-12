package muick

import (
	"fmt"
	"log"
	"net"

	"github.com/adibfahimi/muick/parser"
	"github.com/adibfahimi/muick/types"
)

type Route interface {
	Get(path string, handler func(c *Ctx) error)
	Post(path string, handler func(c *Ctx) error)
	Put(path string, handler func(c *Ctx) error)
	Delete(path string, handler func(c *Ctx) error)

	ListenAndServe(addr ...string)
}

type Server struct {
	Routes map[string]map[string]func(ctx *Ctx) error
	Addr   string
}

func New() *Server {
	return &Server{
		Routes: make(map[string]map[string]func(ctx *Ctx) error),
		Addr:   ":8080",
	}
}

func (s *Server) Get(path string, handler func(c *Ctx) error) {
	if s.Routes["GET"] == nil {
		s.Routes["GET"] = make(map[string]func(ctx *Ctx) error)
	}
	s.Routes["GET"][path] = handler
}

func (s *Server) Post(path string, handler func(c *Ctx) error) {
	if s.Routes["POST"] == nil {
		s.Routes["POST"] = make(map[string]func(ctx *Ctx) error)
	}
	s.Routes["POST"][path] = handler
}

func (s *Server) Put(path string, handler func(c *Ctx) error) {
	if s.Routes["PUT"] == nil {
		s.Routes["PUT"] = make(map[string]func(ctx *Ctx) error)
	}
	s.Routes["PUT"][path] = handler
}

func (s *Server) Delete(path string, handler func(c *Ctx) error) {
	if s.Routes["DELETE"] == nil {
		s.Routes["DELETE"] = make(map[string]func(ctx *Ctx) error)
	}
	s.Routes["DELETE"][path] = handler
}

func HandleConnection(c net.Conn, server *Server) {
	defer c.Close()
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	req := parser.ParseRequest(string(buf[:n]))

	ctx := &Ctx{
		Request: req,
		Response: types.HttpResponse{
			Headers: map[string]string{},
			Version: "1.1",
			Reason:  "OK",
			Body:    "",
			Status:  200,
		},
	}

	handler, ok := server.Routes[req.Method][req.Path]
	if !ok {
		c.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}

	err = handler(ctx)
	if err != nil {
		log.Printf("Error handling request: %v", err)
		c.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}

	c.Write([]byte(ctx.Response.String()))
}

func (s *Server) Listen(addr ...string) {
	if len(addr) > 0 {
		s.Addr = addr[0]
	}

	conn, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("Listening on %s\n", s.Addr)
	for {
		c, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConnection(c, s)
	}
}
