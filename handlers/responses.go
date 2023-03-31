package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type jsonResponse map[string]interface{}

func postError(_w http.ResponseWriter, _errCode int) {
	http.Error(_w, http.StatusText(_errCode), _errCode)

}

func postBodyResponse(_w http.ResponseWriter, _statCode int, _content jsonResponse) {
	if _content != nil {
		js, err := json.Marshal(_content)
		if err != nil {
			postError(_w, http.StatusInternalServerError)
			return
		}
		_w.Header().Set("Content-Type", "application/json")
		_w.WriteHeader(_statCode)
		_w.Write(js)
		return
	}
	_w.WriteHeader(_statCode)
	_w.Write([]byte(http.StatusText(_statCode)))
}

func postOptionsResponse(_w http.ResponseWriter, _methods []string, _content jsonResponse) {
	_w.Header().Set("Allow", strings.Join(_methods, ","))
	postBodyResponse(_w, http.StatusOK, _content)
}
