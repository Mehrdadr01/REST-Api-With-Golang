package handlers

import (
	"net/http"
	"users"
)

func usersGetAll(_w http.ResponseWriter, _r *http.Request) {
	users, err := users.All()
	if err != nil {
		postError(_w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(_w, http.StatusOK, jsonResponse{"users": users})
}
