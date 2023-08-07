package auth

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mikerybka/apps/pkg/web/types"
	"github.com/mikerybka/apps/pkg/web/util"
)

type SessionStore struct {
	Dir string
}

func (store *SessionStore) Create(userID types.ID) (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	var token string
	for {
		token = hex.EncodeToString(randomBytes)
		filename := filepath.Join(store.Dir, token+".json")
		if !util.FileExists(filename) {
			break
		}
	}
	session := types.Session{UserID: userID}
	err = store.Put(token, session)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (store *SessionStore) Put(token string, session types.Session) error {
	b, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(store.Dir, token+".json"), b, 0644)
	if err != nil {
		return err
	}
	return nil
}
