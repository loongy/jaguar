package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/loongy/jaguar-template/actions"
	"github.com/loongy/jaguar-template/models"
)

func PostUser(ctx actions.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Decode http request into model
		user := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error"))
			return
		}

		// Trigger an action to create a user
		user, err := actions.CreateUser(ctx, user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error"))
			return
		}

		// Write the created user to the response
		w.WriteHeader(201)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error"))
		}
	})
}

func GetUsers(ctx actions.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func GetUser(ctx actions.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func PutUser(ctx actions.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func PatchUser(ctx actions.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func DeleteUser(ctx actions.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
