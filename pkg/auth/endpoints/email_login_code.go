package auth

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mikerybka/webmachine/pkg/auth"
	"github.com/mikerybka/webmachine/pkg/email"
	"github.com/mikerybka/webmachine/pkg/web"
)

func EmailLoginCode() {
	var addr string
	var redirect string
	flag.StringVar(&addr, "addr", "", "")
	flag.StringVar(&redirect, "redirect", "", "")
	flag.Parse()

	user := auth.FindUserByEmail(addr)
	if user == nil {
		return
	}

	u, _ := url.Parse(redirect)
	host := u.Host
	code := auth.CreateOTP(host, user.ID)

	subject := fmt.Sprintf("Your OTP for %s", host)
	message := fmt.Sprintf("Your one time password is:\n\n%s\n\nThis code is valid for 15 minutes.", code)
	err := email.Send(addr, subject, message)
	if err != nil {
		web.Return(http.StatusBadRequest, err.Error())
		return
	}
}
