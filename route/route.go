package route

import (
	"github.com/labstack/echo/v4"

	"demonstrasi/constants"
	"demonstrasi/controllers"
	m "demonstrasi/middlewares"

	jwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LoggerMiddleware(e)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	e.POST("/comment", controllers.CreateComment)
	e.GET("/homepage", controllers.GetHomePage)
	e.GET("/galery", controllers.GetGalery)
	e.GET("/monument", controllers.GetMonument)
	e.GET("/commented", controllers.GetComment)

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
	commentAdmin.PUT("/:id", controllers.ValideteComment)

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
