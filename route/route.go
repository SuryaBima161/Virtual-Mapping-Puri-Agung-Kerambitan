package route

import (
	"html/template"
	"io"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demonstrasi/constants"
	"demonstrasi/controllers"
	m "demonstrasi/middlewares"

	jwt "github.com/labstack/echo-jwt/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func New() *echo.Echo {
	e := echo.New()

	m.LoggerMiddleware(e)
	e.Use(middleware.Recover())

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://v8ftmpnc-3000.asse.devtunnels.ms",
			"https://v57q9chz-3000.asse.devtunnels.ms",
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	// ga perlu login
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	e.POST("/comment", controllers.CreateComment)
	e.GET("/homepage", controllers.GetHomePage)
	e.GET("/galery", controllers.GetGalery)
	e.GET("/monument", controllers.GetMonument)
	e.GET("/comment", controllers.GetCommentValidated)

	e.Pre(middleware.RemoveTrailingSlash())

	// ---
	admin := e.Group("/admin", jwt.JWT([]byte(constants.SECRET_KEY)))

	inf := admin.Group("/information")
	inf.POST("", controllers.CreateInformation)
	inf.GET("", controllers.GetInformation)
	inf.GET("/:id", controllers.GetInformationById)
	inf.DELETE("/:id", controllers.DeleteInformation)
	inf.PUT("/:id", controllers.UpdateInfortmation)

	monument := admin.Group("/monument", jwt.JWT([]byte(constants.SECRET_KEY)))
	monument.POST("", controllers.CreateMonument)
	monument.GET("", controllers.GetMonument)
	monument.GET("/:id", controllers.GetMonumentById)
	monument.DELETE("/:id", controllers.DeleteMonument)
	monument.PUT("/:id", controllers.UpdateMonument)

	commentAdmin := admin.Group("/comment")
	commentAdmin.DELETE("/:id", controllers.DeleteComment)
	commentAdmin.GET("", controllers.GetComment)
	commentAdmin.POST("", controllers.CreateComment)
	commentAdmin.PUT("/:id", controllers.UpdateReplyComment)
	commentAdmin.PUT("/validate/:id", controllers.ValidateComment)

	galeryAdmin := admin.Group("/galery")
	galeryAdmin.POST("", controllers.CreateGalery)
	galeryAdmin.GET("", controllers.GetGalery)
	galeryAdmin.GET("/:id", controllers.GetCommentByIdGalery)
	galeryAdmin.DELETE("/:id", controllers.DeleteGalery)
	galeryAdmin.PUT("/:id", controllers.UpdateGalery)

	user := e.Group("/user")

	commentUser := user.Group("/comment")
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/comment.html")),
	}
	e.Renderer = renderer

	// Serve captcha images
	commentUser.GET("/captcha/:id", func(c echo.Context) error {
		captchaId := c.Param("id")
		if captchaId == "" {
			return c.NoContent(http.StatusBadRequest)
		}
		captcha.WriteImage(c.Response().Writer, captchaId, captcha.StdWidth, captcha.StdHeight)
		return nil
	})

	// Serve comment page with captcha
	commentUser.GET("", func(c echo.Context) error {
		captchaId := captcha.New()
		return c.Render(http.StatusOK, "comment.html", map[string]interface{}{
			"CaptchaID": captchaId,
		})
	})

	// Comment route with captcha middleware
	commentUser.POST("", m.Captcha(controllers.CreateComment))

	return e
}
