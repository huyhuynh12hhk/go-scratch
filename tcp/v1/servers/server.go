package servers

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Server include all behavior op a TCP server
type Server struct {
	port     int
	listener net.Listener
	Router   *Router
}

// Initialization for a Server
func NewServer(port int) Server {
	return Server{
		port:   port,
		Router: NewRouter(),
	}
}

// Run server after setup behaviors for server
func (s *Server) Run() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalln(err)
	}

	defer ln.Close()
	s.listener = ln

	// start listening
	s.startListening()
}


// Accept and listen a connect
func (s *Server) startListening() {
	log.Printf("Server start listening on port %d\n", s.port)
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("Accept error: ", err)
			continue
		}
		log.Println("New connection to the server: ", conn.RemoteAddr())
		// For multiplex request
		go s.handleConnect(conn)
	}
}

// Handle the request in income of connection
func (s *Server) handleConnect(conn net.Conn) {
	ctx := NewContext(&conn)
	defer conn.Close()

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		// Print the reading http body
		log.Println(ln)

		// Processing request
		if i == 0 {
			// Use router to route the request and process
			s.Router.ProcessRequest(ctx, ln)
			// Process full of request
		}
		if ln == "" {
			// Done header
			break
		}
		i++
	}
}
