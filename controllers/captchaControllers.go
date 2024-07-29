package controllers

import (
	"net/http"

	"github.com/dchest/captcha"
)

func Captcha(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		captchaId := r.FormValue("captchaId")
		captchaSolution := r.FormValue("captchaSolution")
		if captcha.VerifyString(captchaId, captchaSolution) {
			w.Write([]byte("Captcha verified successfully!"))
		} else {
			w.Write([]byte("Captcha verification failed."))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
