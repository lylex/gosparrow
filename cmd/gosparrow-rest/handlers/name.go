package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"gosparrow/cmd/gosparrow-rest/app"
)

// GetName used to retrieve the name of the app.
func GetName(session *app.Session, w http.ResponseWriter, r *http.Request) (
	status int, data interface{}, err error) {

	status = http.StatusInternalServerError

	prefix := mux.Vars(r)["prefix"]

	return http.StatusOK, map[string]string{
		"name": prefix + "gosparrow",
	}, nil
}
