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
			usersPostOne(_writer, _req)
			return
		case http.MethodHead:
			usersGetAll(_writer, _req)
			return
		case http.MethodOptions:
			postOptionsResponse(_writer, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
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

	id := bson.ObjectIdHex(path)

	switch _req.Method {
	case http.MethodGet:
		usersGetOne(_writer, _req, id)
		return
	case http.MethodPut:
		usersPutOne(_writer, _req, id)
		return
	case http.MethodPatch:
		usersPathOne(_writer, _req, id)
		return
	case http.MethodDelete:
		usersDeleteOne(_writer, _req, id)
		return
	case http.MethodHead:
		usersGetOne(_writer, _req, id)
		return
	case http.MethodOptions:
		postOptionsResponse(_writer, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(_writer, http.StatusMethodNotAllowed)
		return
	}
}
