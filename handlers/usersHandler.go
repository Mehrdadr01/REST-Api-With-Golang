package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"users"

	"encoding/json"

	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

//  errors //

var (
	ErrBodyRequestEmpty = errors.New("request body empty")
	ErrUserRequired     = errors.New("a user is required")
)

func bodyToUser(_r *http.Request, _usr *users.User) error {
	if _r.Body == nil {
		return ErrBodyRequestEmpty
	}
	if _usr == nil {
		return ErrUserRequired
	}
	bdy, err := ioutil.ReadAll(_r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bdy, _usr)
}

/////////////////////////////////////////////////////////////////
func usersGetAll(_w http.ResponseWriter, _r *http.Request) {
	users, err := users.All()
	if err != nil {
		postError(_w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(_w, http.StatusOK, jsonResponse{"users": users})
}

/////////////////////////////////////////////////////////////////
func usersPostOne(_w http.ResponseWriter, _r *http.Request) {
	usr := new(users.User)
	err := bodyToUser(_r, usr)
	if err != nil {
		postError(_w, http.StatusBadRequest)
		return
	}
	usr.ID = bson.NewObjectId()
	err = usr.Save()
	// we get 2 errors : 1.for database and 2.failed validation
	if err != nil {
		if err == users.ErrRecordInvalid {
			postError(_w, http.StatusBadRequest)
		} else {
			postError(_w, http.StatusInternalServerError)
		}
		return
	}
	_w.Header().Set("Location", "/users/"+usr.ID.Hex())
	_w.WriteHeader(http.StatusCreated)
}

///////////////////////////////////////////////////////////////
func usersGetOne(_w http.ResponseWriter, _ *http.Request, _id bson.ObjectId) {
	usr, err := users.One(_id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(_w, http.StatusNotFound)
			return
		}
		postError(_w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(_w, http.StatusOK, jsonResponse{"user": usr})
}
