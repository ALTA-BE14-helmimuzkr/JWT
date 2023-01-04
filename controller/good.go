package controller

import (
	"api/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GoodController struct {
	Mdl    *model.GoodModel
	JWTKey string
}

func (gc *GoodController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		good := model.Good{}
		if err := c.Bind(&good); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		userID := ExtractToken(c)
		good.UserID = userID

		good, err := gc.Mdl.Insert(good)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    good,
			"message": "sukses menambahkan data",
		})
	}
}

func (gc *GoodController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ExtractToken(c)
		goods, err := gc.Mdl.GetAll(userID)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    goods,
			"message": "sukses mendapatkan semua data"})
	}
}
