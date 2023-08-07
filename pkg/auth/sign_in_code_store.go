package auth

import (
	"github.com/mikerybka/apps/pkg/web/types"
	"github.com/mikerybka/apps/pkg/web/util"
)

type SignInCodeStore struct {
	Dir string
}

func (store *SignInCodeStore) Create(userID types.ID) (string, error) {
	code := types.SignInCode{Code: util.NewSixDigitCode(), UserID: userID}
	for {
		if !util.FileExists(util.JSONFileName(store.Dir, code.Code)) {
			break
		}
		code.Code = util.NewSixDigitCode()
	}
	err := store.Put(code.Code, code)
	if err != nil {
		return "", err
	}
	return code.Code, nil
}

func (store *SignInCodeStore) Get(code string) (*types.SignInCode, error) {
	var sc types.SignInCode
	err := util.ReadJSON(util.JSONFileName(store.Dir, code), &sc)
	if err != nil {
		return nil, err
	}
	return &sc, nil
}

func (store *SignInCodeStore) Put(code string, sc types.SignInCode) error {
	return util.WriteJSON(util.JSONFileName(store.Dir, code), sc)
}

func (store *SignInCodeStore) Delete(code string) error {
	return util.DeleteFile(util.JSONFileName(store.Dir, code))
}

func (store *SignInCodeStore) DeleteAll() error {
	return util.DeleteAllFiles(store.Dir)
}
