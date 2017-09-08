package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/loongy/jaguar/actions"
	"github.com/loongy/jaguar/api/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler handlers.Handler
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
			Handler: handlers.Health,
		},
		Route{
			Method:  "POST",
			Path:    "/users",
			Handler: handlers.PostUser,
		},
		Route{
			Method:  "GET",
			Path:    "/users",
			Handler: handlers.GetUsers,
		},
		Route{
			Method:  "GET",
			Path:    "/users/{userID}",
			Handler: handlers.GetUser,
		},
		Route{
			Method:  "PUT",
			Path:    "/users/{userID}",
			Handler: handlers.PutUser,
		},
		Route{
			Method:  "PATCH",
			Path:    "/users/{userID}",
			Handler: handlers.PatchUser,
		},
		Route{
			Method:  "DELETE",
			Path:    "/users/{userID}",
			Handler: handlers.DeleteUser,
		},
		Route{
			Method:  "GET",
			Path:    "/me",
			Handler: handlers.GetUser,
		},
		Route{
			Method:  "PUT",
			Path:    "/me",
			Handler: handlers.PutUser,
		},
		Route{
			Method:  "PATCH",
			Path:    "/me",
			Handler: handlers.PatchUser,
		},
		Route{
			Method:  "DELETE",
			Path:    "/me",
			Handler: handlers.DeleteUser,
		},
	}

	r := &Server{mux.NewRouter()}
	for _, route := range routes {
		r.Methods(route.Method).Path(route.Path).Handler(
			handlers.Logging(os.Stdout)(handlers.Recovery(route.Handler))(ctx),
		)
	}
	return r, nil
}

func (server *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, server)
}
