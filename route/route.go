package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demonstrasi/constants"
	"demonstrasi/controllers"
	m "demonstrasi/middlewares"

	jwt "github.com/labstack/echo-jwt/v4"
)

func New() *echo.Echo {
	e := echo.New()

	m.LoggerMiddleware(e)
	e.Use(middleware.Recover())

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",                     // Ganti dengan origin frontend Anda
			"https://v8ftmpnc-3000.asse.devtunnels.ms/", // Contoh: URL publik yang diizinkan
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Pre(middleware.RemoveTrailingSlash())
	//ga perlu login
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	e.POST("/comment", controllers.CreateComment)
	e.GET("/homepage", controllers.GetHomePage)
	e.GET("/galery", controllers.GetGalery)
	e.GET("/monument", controllers.GetMonument)
	e.GET("/comment", controllers.GetCommentValidated)
	//---
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
	commentAdmin.PUT("validate/:id", controllers.ValideteComment)

	galeryAdmin := admin.Group("/galery")
	galeryAdmin.POST("", controllers.CreateGalery)
	galeryAdmin.GET("", controllers.GetGalery)
	galeryAdmin.DELETE("/:id", controllers.DeleteGalery)
	galeryAdmin.PUT("/:id", controllers.UpdateGalery)

	user := e.Group("/user")

	commentUser := user.Group("/comment")
	commentUser.POST("", controllers.CreateComment)
	return e
}
