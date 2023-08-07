package auth

import (
	"net/url"

	"github.com/mikerybka/apps/pkg/web"
)

func Authorize() {
	r := web.Request("/auth/authorize", "POST")
	r.ParseForm()
	token, err := r.Cookie("token")
	if err != nil {
		web.Return(401, "unauthorized")
	}
	redirect := r.FormValue("redirect")
	if redirect == "" {
		web.Return(400, "missing redirect")
	}
	redirectURL, err := url.Parse(redirect)
	if err != nil {
		web.Return(400, "invalid redirect")
	}
	query := redirectURL.Query()
	query.Set("token", token.Value)
	redirectURL.RawQuery = query.Encode()
	web.Redirect(redirectURL.String())
}
