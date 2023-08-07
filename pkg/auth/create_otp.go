package auth

import (
	"fmt"
	"time"

	"github.com/mikerybka/webmachine/pkg/types"
	"github.com/mikerybka/webmachine/pkg/util"
	"github.com/mikerybka/webmachine/pkg/web"
)

func CreateOTP(host string, userID types.ID) string {
	code := util.NewSixDigitCode()
	path := fmt.Sprintf("/data/auth/otps/%s.json", code)
	for util.FileExists(path) {
		code = util.NewSixDigitCode()
		path = fmt.Sprintf("/data/auth/otps/%s.json", code)
	}

	err := util.WriteJSON(path, userID)
	if err != nil {
		return ""
	}

	removeAt := time.Now().Add(15 * time.Minute)
	web.ScheduleTask(removeAt, "auth.RemoveOTP", code)

	return code
}
