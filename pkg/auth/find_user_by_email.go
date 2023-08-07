package auth

import (
	"fmt"

	"github.com/mikerybka/webmachine/pkg/types"
	"github.com/mikerybka/webmachine/pkg/util"
)

func FindUserByEmail(email string) *types.User {
	index := emailIndex()
	userID, ok := index[email]
	if !ok {
		return nil
	}
	path := fmt.Sprintf("/data/users/%s.json", userID)
	var user types.User
	util.ReadJSON(path, &user)
	return &user
}
