package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/loongy/jaguar/actions"
	"github.com/loongy/jaguar/models"
)

func PostUser(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {

		// Decode http request into model
		user := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			return err
		}

		// Trigger an action to create a user
		user, err := actions.CreateUser(ctx, user)
		if err != nil {
			return err
		}

		// Write the created user to the response
		w.WriteHeader(201)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			return err
		}

		return nil
	}
}

func GetUsers(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func GetUser(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func PutUser(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func PatchUser(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func DeleteUser(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func GetMe(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func PutMe(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func PatchMe(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func DeleteMe(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}
