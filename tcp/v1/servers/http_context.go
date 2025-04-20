package servers

import (
	"fmt"
	"net"

	"github.com/google/uuid"
)

// HttpContext contains states of a new request
type HttpContext struct {
	id   string
	conn net.Conn
}

// Initialization for a HttpContext
func NewContext(conn *net.Conn) *HttpContext {
	id := uuid.New()
	return &HttpContext{
		id:   id.String(),
		conn: *conn,
	}
}

// Mutate response header
func (ctx *HttpContext) WriteHeader(statusCode int) {
	panic("WriteHeader function not implement yet.")
}

// Compose response and return to client
func (ctx *HttpContext) WriteResponse(data []byte) {
	contentType := "application/json"

	// TODO: Implement json convert or similar thing
	body := string(data)

	// TODO: Implement status code mapping
	fmt.Fprint(ctx.conn, "HTTP/1.1 200 OK\r\n")
	// TODO: Implement common header mapping
	fmt.Fprintf(ctx.conn, "Content-Length: %d\r\n", len(body))
	// TODO: Implement content types mapping
	fmt.Fprintf(ctx.conn, "Content-Type: %s\r\n", contentType)
	fmt.Fprint(ctx.conn, "\r\n")
	fmt.Fprint(ctx.conn, body)
}

// Simple response way
func (ctx *HttpContext) WriteSimpleResponse() {
	
	body := `
	<html>
	<head>
	<title>Example Simple</title>
	</head>
	<body>
	<strong>
	Hello World
	</strong>
	</body>
	</html>
	`

	fmt.Fprint(ctx.conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(ctx.conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(ctx.conn, "Content-Type: text/html\r\n")
	fmt.Fprint(ctx.conn, "\r\n")
	fmt.Fprint(ctx.conn, body)
}

