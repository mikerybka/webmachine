package auth

import (
	"path/filepath"
)

type DB struct {
	Dir string
}

func NewDB(dir string) *DB {
	return &DB{Dir: dir}
}

func (db *DB) Users() *UserStore {
	return &UserStore{Dir: filepath.Join(db.Dir, "users")}
}

func (db *DB) Sessions() *SessionStore {
	return &SessionStore{Dir: filepath.Join(db.Dir, "sessions")}
}

func (db *DB) SignInCodes() *SignInCodeStore {
	return &SignInCodeStore{Dir: filepath.Join(db.Dir, "sign_in_codes")}
}
