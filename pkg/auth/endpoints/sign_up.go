package auth

import (
	"net/http"

	"github.com/mikerybka/apps/pkg/web"
)

func SignUp() {
	r := web.Request("/auth/sign-up", "POST")
	r.ParseForm()
	firstName := r.FormValue("first-name")
	lastName := r.FormValue("last-name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	db := web.NewDB("/data")

	_, err := db.Users().Create(
		firstName,
		lastName,
		email,
		phone,
	)
	if err != nil {
		web.Return(http.StatusBadRequest, err.Error())
	}

	web.Redirect("/auth/sign-in")
}
