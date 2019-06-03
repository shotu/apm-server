package custom_server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type CustomServer struct {
	router *mux.Router
}

func NewServer() *CustomServer {
	s := CustomServer{router: mux.NewRouter()}
	NewUserRouter(s.newSubrouter("/custom_end_point"))
	return &s
}

func (s *CustomServer) Start() {
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *CustomServer) newSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
