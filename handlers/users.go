package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/hitesh-goel/loomx/internal/platform/db"
	"github.com/hitesh-goel/loomx/internal/platform/web"
	"github.com/pkg/errors"
)

// User represents the User API method handler set.
type User struct {
	MasterDB *db.DB
}

// NewUser create a new user
type NewUser struct {
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

// UserDetails create a new user
type UserDetails struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

//Create a new User
func (u *User) Create(ctx context.Context, log *log.Logger, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	var newU NewUser

	if err := web.Unmarshal(r.Body, &newU); err != nil {
		return errors.Wrap(err, "")
	}

	id, err := uuid.NewV4()

	if err != nil {
		return err
	}

	userDetails := UserDetails{
		UserName: newU.UserName,
		Email:    newU.Email,
		ID:       id.String(),
	}

	encodedData, err := json.Marshal(userDetails)
	if err != nil {
		return err
	}

	dbConn := u.MasterDB

	// Save User by ID
	fID := func(db *leveldb.DB) error {
		// TODO: while inserting validate if email already exists by checking with email index
		return db.Put([]byte(userDetails.ID), encodedData, nil)
	}

	// Save User by Email
	fEmail := func(db *leveldb.DB) error {
		// TODO: while inserting validate if already exists or not
		return db.Put([]byte(userDetails.Email), encodedData, nil)
	}

	if err := dbConn.Execute(fID); err != nil {
		return errors.Wrap(err, "")
	}

	if err := dbConn.Execute(fEmail); err != nil {
		// TODO: if error remove the data from ID as well and throw error
		return errors.Wrap(err, "")
	}

	status := struct {
		ID string `json:"id"`
	}{
		ID: userDetails.ID,
	}
	web.Respond(ctx, log, w, status, http.StatusCreated)
	return nil
}

//Retrieve an existing User
func (u *User) Retrieve(ctx context.Context, log *log.Logger, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	queryParams := r.URL.Query()
	var email, id string
	var usr *UserDetails
	if len(queryParams["email"]) > 0 {
		// TODO: add email validation
		email = queryParams["email"][0]
	}
	dbConn := u.MasterDB
	if params["id"] != "" {
		id = params["id"]
	}
	if email != "" {
		id = email
	}
	usr, err := query(dbConn, id)
	if err != nil {
		return err
	}
	web.Respond(ctx, log, w, usr, http.StatusOK)
	return nil
}

// Query by ID or Email
func query(dbConn *db.DB, id string) (*UserDetails, error) {
	var user *UserDetails
	f := func(db *leveldb.DB) error {
		data, err := db.Get([]byte(id), nil)
		if err != nil {
			return err
		}
		var u UserDetails
		if err := json.Unmarshal(data, &u); err != nil {
			return errors.Wrap(err, "")
		}
		user = &u
		return nil
	}
	if err := dbConn.Execute(f); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return user, nil
}
