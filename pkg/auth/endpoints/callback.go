package auth

import "github.com/mikerybka/apps/pkg/web"

func Callback() {
	r := web.Request("/auth/callback", "GET")
	r.ParseForm()
	token := r.FormValue("token")
	redirect := r.FormValue("redirect")

	web.SetCookie("token", token)
	web.Redirect(redirect)
}
