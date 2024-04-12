package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type typeOfRequest struct {
	Method, Path string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[typeOfRequest]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[typeOfRequest]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers[typeOfRequest{
		Method: method,
		Path:   path,
	}] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for request, handler := range s.Handlers {

		switch request.Method {
		case "GET":
			s.Router.Get(request.Path, handler)
		case "POST":
			s.Router.Post(request.Path, handler)
		// Adicione outros métodos aqui, se necessário
		default:
			panic("Método HTTP não suportado: " + request.Method)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
