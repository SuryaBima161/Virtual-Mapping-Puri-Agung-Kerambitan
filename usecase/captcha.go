package usecase

import (
	"net/http"

	"github.com/dchest/captcha"
)

func FormCaptcha(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Captcha Example</title>
	</head>
	<body>
		<form action="/submit" method="POST">
			<img src="/captcha?id=` + captcha.New() + `" alt="CAPTCHA">
			<input type="text" name="captchaSolution">
			<input type="hidden" name="captchaId" value="` + captcha.New() + `">
			<input type="submit" value="Submit">
		</form>
	</body>
	</html>`
	w.Write([]byte(html))
}


