package util

import "github.com/mikerybka/apps/pkg/web/types"

type JSONStore[T any] struct {
	Dir string
}

func (store *JSONStore[T]) Get(id types.ID) (*T, error) {
	var v T
	path := JSONFileName(store.Dir, id.String())
	err := ReadJSON(path, &v)
	return &v, err
}

func (store *JSONStore[T]) Put(id types.ID, v T) error {
	path := JSONFileName(store.Dir, id.String())
	return WriteJSON(path, v)
}

func (store *JSONStore[T]) Delete(id types.ID) error {
	path := JSONFileName(store.Dir, id.String())
	return DeleteFile(path)
}

func (store *JSONStore[T]) Index(name string) *Index {
	path := JSONFileName(store.Dir, name)
	return &Index{File: path}
}
