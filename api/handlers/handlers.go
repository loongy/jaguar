package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/loongy/jaguar/actions"
)

// HandlerFunc is an http.Handler that can also return an error.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// NewHandlerFunc wraps an http.Handler and always returns a nil error.
func NewHandlerFunc(h http.Handler) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		h.ServeHTTP(w, r)
		return nil
	}
}

// ServeHTTP implements http.Handler by panicking if the returned error is
// anything other than nil.
func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		panic(err)
	}
}

// Handler is a function that accepts an actions.Context and returns a
// HandlerFunc. This allows a Handler to be used to create a HandlerFunc
// that has access to an actions.Context without the HandlerFunc needing to
// change its type signature.
type Handler func(actions.Context) HandlerFunc

// Middleware wraps a Handler in an outer Handler and returns the outer
// Handler. This allows the Middleware to intercept requests and perform
// pre-processing and post-processing.
type Middleware func(Handler) Handler

// Recovery is a Middleware that wraps an http.Handler and will recover from
// any panic that happens in that http.Handler. This stop the server crashing
// due to an unexpected panic.
func Recovery(next Handler) Handler {
	return func(ctx actions.Context) HandlerFunc {
		return NewHandlerFunc(handlers.RecoveryHandler()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next(ctx).ServeHTTP(w, r)
		})))
	}
}

// Logging returns Middleware that wraps a Handler and logs its behaviour
// using the Apache Common Log Format. The logs will be written to the
// provided io.Writer.
func Logging(w io.Writer) Middleware {
	return func(next Handler) Handler {
		return func(ctx actions.Context) HandlerFunc {
			return NewHandlerFunc(handlers.LoggingHandler(w, next(ctx)))
		}
	}
}

// SessionDecoder exposes the functionality to decode sessions from an
// http.Request and write them to an actions.Context.
type SessionDecoder interface {
	Decode(ctx actions.Context, r *http.Request) (actions.Context, error)
}

// SessionEncoder exposes the functionality to encode the session details of
// an actions.Context into an http.ResponseWriter.
type SessionEncoder interface {
	Encode(ctx actions.Context, w http.ResponseWriter) error
}

// Session returns Middleware that extracts the current authentication
// models.Session from the request, and embeds that information into an
// actions.Context. This is used to call the original Handler, thus giving it
// access to the models.Session of the request.
func Session(decoder SessionDecoder, encoder SessionEncoder) Middleware {
	return func(next Handler) Handler {
		return func(ctx actions.Context) HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) error {
				// Decode the session into a new actions.Context.
				ctx, err := decoder.Decode(ctx, r)
				if err != nil {
					return err
				}
				// Use this context to create the next handler.
				if err := next(ctx)(w, r); err != nil {
					return err
				}
				// Use this context to encode the new session.
				if err := encoder.Encode(ctx, w); err != nil {
					return err
				}
				return nil
			}
		}
	}
}
