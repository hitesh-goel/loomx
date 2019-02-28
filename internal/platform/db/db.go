package db

import (
	"log"

	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

// ErrInvalidDBProvided is returned in the event that an uninitialized db is
// used to perform actions against.
var ErrInvalidDBProvided = errors.New("invalid DB provided")

// DB is a collection of support for different DB technologies
type DB struct {

	// LevelDB Support.
	database *leveldb.DB
}

// New returns a new DB value for use with LevelDB based on a registered
func New(path string) (*DB, error) {
	ldb, err := leveldb.OpenFile(path, nil)
	db := DB{
		database: ldb,
	}
	if err != nil {
		log.Fatal("Yikes!")
	}
	return &db, nil
}

// Close closes a DB value being used with MongoDB.
func (db *DB) Close() {
	db.database.Close()
}

// Execute is used to execute MongoDB commands.
func (db *DB) Execute(f func(*leveldb.DB) error) error {
	if db == nil {
		return errors.Wrap(ErrInvalidDBProvided, "db == nil || db.session == nil")
	}

	return f(db.database)
}
