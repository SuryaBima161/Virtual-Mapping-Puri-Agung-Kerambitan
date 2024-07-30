package middlewares

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
)

// Captcha middleware function to verify captcha before proceeding to the next handler
func Captcha(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		captchaId := c.FormValue("captchaId")
		captchaSolution := c.FormValue("captchaSolution")

		if captchaId == "" || captchaSolution == "" || !captcha.VerifyString(captchaId, captchaSolution) {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "Captcha verification failed",
			})
		}

		// If captcha verification is successful, proceed to the next handler
		return next(c)
	}
}
