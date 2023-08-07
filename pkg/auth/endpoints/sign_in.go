package auth

import (
	"flag"
	"net/http"

	"github.com/mikerybka/apps/pkg/web"
)

func SignIn() {
	var code string
	var redirect string
	flag.StringVar(&code, "code", "", "")
	flag.StringVar(&redirect, "redirect", "", "")
	flag.Parse()

	db := web.NewDB("/data")

	signInCode, err := db.SignInCodes().Get(code)
	if err != nil {
		web.Return(http.StatusBadRequest, "invalid sign in code")
	}

	user, err := db.Users().Get(signInCode.UserID)
	if err != nil {
		web.Report(err)
		web.Return(http.StatusInternalServerError, "user not found")
	}

	token, err := db.Sessions().Create(user.ID)
	if err != nil {
		web.Return(http.StatusInternalServerError, err.Error())
	}

	web.SetCookie("token", token)
	web.Redirect(redirect)
}
