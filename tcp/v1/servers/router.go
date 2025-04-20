package servers

import (
	
	"fmt"
	"log"
	"strings"
)

// Router to manage request pattern registration 
// and then route to process these registered patterns
type Router struct {
	routes map[string]func(*HttpContext)
}

// Initialization for a Router
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]func(*HttpContext)),
	}
}

// Register new pattern with handler function
func (r *Router) AddRoute(pattern string, handler func(*HttpContext)) {
	// TODO: Handle by http method
	r.routes[pattern] = handler
	
}

// Route request to right handler
func (r *Router) ProcessRequest(ctx *HttpContext, path string) error {

	// method
	method := strings.Fields(path)[0]
	// uri
	uri := strings.Fields(path)[1]
	log.Println("-> METHOD", method)
	log.Println("-> PATH", uri)

	handler, ok := r.routes[uri]
	if !ok || handler == nil {
		log.Printf("Error when try to processing path %s\n", uri)
		return fmt.Errorf("error occur when try to handle path %s",uri)
	}

	handler(ctx)

	return nil
}
