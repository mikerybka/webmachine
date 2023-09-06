package util

import (
	"errors"
	"os"

	"github.com/mikerybka/webmachine/pkg/types"
)

type Index struct {
	File string
}

func (i *Index) Get(key string) (types.ID, bool) {
	index := map[string]types.ID{}
	err := ReadJSON(i.File, &index)
	if err != nil {
		return 0, false
	}
	val, ok := index[key]
	return val, ok
}

func (i *Index) Set(key string, val types.ID) error {
	index := map[string]types.ID{}
	err := ReadJSON(i.File, &index)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	index[key] = val
	err = WriteJSON(i.File, index)
	if err != nil {
		return err
	}
	return nil
}
