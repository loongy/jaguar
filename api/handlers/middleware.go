package handlers

import "net/http"
import "github.com/gorilla/handlers"

type Middleware func(http.Handler) http.Handler

var RecoveryHandler = handlers.RecoveryHandler
var LoggingHandler = handlers.LoggingHandler
