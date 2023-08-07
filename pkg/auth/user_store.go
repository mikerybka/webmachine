package auth

import (
	"fmt"
	"strings"

	"github.com/mikerybka/apps/pkg/web/types"
	"github.com/mikerybka/apps/pkg/web/util"
)

type UserStore struct {
	Dir string
}

func (us *UserStore) Get(id types.ID) (*types.User, error) {
	store := util.JSONStore[types.User]{Dir: us.Dir}
	return store.Get(id)
}

func (us *UserStore) Create(
	firstName string,
	lastName string,
	email string,
	phone string,
) (*types.User, error) {
	store := util.JSONStore[types.User]{Dir: us.Dir}

	// Normalize input
	firstName = util.NormalizeName(firstName)
	lastName = util.NormalizeName(lastName)
	email = strings.TrimSpace(strings.ToLower(email))
	phone = util.FilterNonDigits(phone)

	// Check for existing email or phone
	_, ok := store.Index("emails").Get(email)
	if ok {
		return nil, fmt.Errorf("email %s already registered", email)
	}
	_, ok = store.Index("phones").Get(phone)
	if ok {
		return nil, fmt.Errorf("phone %s already registered", phone)
	}

	// Create the new user
	user := types.User{
		ID:        us.NextID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	// Write to storage
	err := store.Put(user.ID, user)
	if err != nil {
		return nil, err
	}
	err = store.Index("emails").Set(email, user.ID)
	if err != nil {
		return nil, err
	}
	err = store.Index("phones").Set(phone, user.ID)
	if err != nil {
		return nil, err
	}

	// Return the new user
	return &user, nil
}

func (us *UserStore) NextID() types.ID {
	var id types.ID
	path := util.JSONFileName(us.Dir, "nextid")
	util.ReadJSON(path, &id)
	nextID := id + 1
	err := util.WriteJSON(path, nextID)
	if err != nil {
		panic(err)
	}
	return id
}
