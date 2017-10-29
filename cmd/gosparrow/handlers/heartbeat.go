package handlers

import (
	"net/http"

	"gosparrow/cmd/gosparrow/app"
)

// Heartbeat used to offer heartbeat check for the service
func Heartbeat(session *app.Session, w http.ResponseWriter, r *http.Request) (
	status int, data interface{}, err error) {

	return http.StatusOK, nil, nil
}
