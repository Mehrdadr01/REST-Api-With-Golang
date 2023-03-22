package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouter(_writer http.ResponseWriter, _req *http.Request) {

	path := strings.TrimSuffix(_req.URL.Path, "/")
	if path == "/users" {
		switch _req.Method {
		case http.MethodGet:
			usersGetAll(_writer, _req)
			return
		case http.MethodPost:
			return
		default:
			postError(_writer, http.StatusMethodNotAllowed)
			return
		}
	}
	// fmt.Println(_req.URL.Path)
	path = strings.TrimPrefix(path, "/users/")
	if !bson.IsObjectIdHex(path) {
		postError(_writer, http.StatusNotFound)
		return
	}

	// id := bson.ObjectIdHex(path)

	switch _req.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(_writer, http.StatusMethodNotAllowed)
		return
	}
}
