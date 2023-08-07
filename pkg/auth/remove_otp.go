package auth

import (
	"fmt"
	"os"
)

func RemoveOTP(code string) {
	path := fmt.Sprintf("/data/auth/otps/%s.json", code)
	os.Remove(path)
}
