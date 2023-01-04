package main

import (
	"api/config"
	"api/controller"
	"api/model"
	"log"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init Echo
	e := echo.New()

	// Setup Database
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	// Setup Model
	userModel := model.UserModel{DB: db}
	goodModel := model.GoodModel{DB: db}

	// Setup Controller
	userController := controller.UserControll{
		Mdl:    &userModel,
		JWTKey: cfg.JWTKEY,
	}
	goodController := controller.GoodController{
		Mdl:    &goodModel,
		JWTKey: cfg.JWTKEY,
	}

	// Setup Middleware
	e.Pre(middleware.RemoveTrailingSlash()) // fungsi ini dijalankan sebelum routing
	e.Use(middleware.CORS())                // WAJIB DIPAKAI agar tidak terjadi masalah permission
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_custom}, ${method}, ${uri}, status=${status}\n", CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	// Route
	e.POST("/register", userController.Insert())
	e.POST("/login", userController.Login())

	user := e.Group("/users", echojwt.JWT([]byte(cfg.JWTKEY)))
	user.GET("/", userController.GetAll())
	user.GET("/profile", userController.GetID())
	user.PUT("/", userController.Update())
	user.DELETE("/", userController.Delete())

	good := e.Group("/goods", echojwt.JWT([]byte(cfg.JWTKEY)))
	good.GET("/", goodController.GetAll())
	good.POST("/", goodController.Create())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
