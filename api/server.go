package api

import (
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/loongy/jaguar/actions"
	"github.com/loongy/jaguar/api/handlers"
)

type Route struct {
	Method     string
	Path       string
	Middleware []handlers.Middleware
	Handler    http.Handler
}

type Routes []Route

type Server struct {
	*mux.Router
}

func NewServer(ctx actions.Context) (*Server, error) {
	routes := Routes{
		Route{
			Method:  "GET",
			Path:    "/health",
			Handler: handlers.Health(),
		},
		Route{
			Method:  "POST",
			Path:    "/users",
			Handler: handlers.PostUser(ctx),
		},
		Route{
			Method:  "GET",
			Path:    "/users",
			Handler: handlers.GetUsers(ctx),
		},
		Route{
			Method:  "GET",
			Path:    "/users/{userID}",
			Handler: handlers.GetUser(ctx),
		},
		Route{
			Method:  "PUT",
			Path:    "/users/{userID}",
			Handler: handlers.PutUser(ctx),
		},
		Route{
			Method:  "PATCH",
			Path:    "/users/{userID}",
			Handler: handlers.PatchUser(ctx),
		},
		Route{
			Method:  "DELETE",
			Path:    "/users/{userID}",
			Handler: handlers.DeleteUser(ctx),
		},
	}

	r := &Server{mux.NewRouter()}
	for _, route := range routes {
		r.Methods(route.Method).Path(route.Path).Handler(handlers.RecoveryHandler()(handlers.LoggingHandler(
			os.Stdout, context.ClearHandler(route.Handler),
		)))
	}
	return r, nil
}

func (server *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, server)
}
