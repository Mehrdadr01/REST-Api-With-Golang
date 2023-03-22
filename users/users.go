package users

import (
	"errors"

	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

// the User structure is storing data for a single user
type User struct {
	ID   bson.ObjectId `json: "id" storm:"id"`
	Name string        `json: "name"`
	Role string        `json: "role"`
}

const (
	dbPATH = "users.db"
)

// errors

var (
	ErrRecordInvalid = errors.New("The record is invalid !")
)

// All returns a all users from the database
func All() ([]User, error) {
	db, err := storm.Open(dbPATH)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// OneUser returns a single user from the database
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPATH)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	singleUser := new(User)
	err = db.One("ID", id, singleUser)
	if err != nil {
		return nil, err
	}
	return singleUser, err
}

// Delete removes an a given record from the database
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPATH)
	if err != nil {
		return err
	}
	defer db.Close()
	singleUser := new(User)
	err = db.One("ID", id, singleUser)
	if err != nil {
		return err
	}
	return db.DeleteStruct(singleUser)

}

// Save create || update an given record in database
func (singleUser *User) Save() error {
	if err := singleUser.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPATH)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(singleUser)
}

// validate checks to make sure our data record is valid

func (singleUser *User) validate() error {
	if singleUser.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
