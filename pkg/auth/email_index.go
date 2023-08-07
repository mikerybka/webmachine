package auth

import "github.com/mikerybka/webmachine/pkg/util"

func emailIndex() map[string]string {
	path := "/data/users/emails.json"
	index := map[string]string{}
	util.ReadJSON(path, &index)
	return index
}
