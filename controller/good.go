package controller

import (
	"api/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GoodController struct {
	Mdl    *model.GoodModel
	JWTKey string
}

// Menambahkan barang baru
func (gc *GoodController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		good := model.Good{}
		if err := c.Bind(&good); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		userID := ExtractToken(c) // id pemilik dari proses extract token
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

// Mendapatkan barang yang dipunyai oleh pemilik
func (gc *GoodController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ExtractToken(c) // id pemilik dari proses extract token
		goods, err := gc.Mdl.GetAll(userID)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    goods,
			"message": "sukses mendapatkan semua data",
		})
	}
}

// Mendapatkan barang punya pemilik yang dipilih menggunakan id barang
func (gc *GoodController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ExtractToken(c)                  // id pemilik dari proses extract token
		goodID, err := strconv.Atoi(c.Param("id")) // id barang dari uri parameter
		if err != nil {
			log.Println("convert id error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gunakan input angka",
			})
		}

		good, err := gc.Mdl.GetByID(userID, goodID)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    good,
			"message": "sukses mendapatkan semua data",
		})
	}
}

// Mendapatkan barang punya pemilik yang dipilih menggunakan id barang
func (gc *GoodController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		good := model.Good{}
		if err := c.Bind(&good); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		userID := ExtractToken(c) // id pemilik dari proses extract token
		good.UserID = userID

		good, err := gc.Mdl.Update(good)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    good,
			"message": "berhasil update data",
		})
	}
}

func (gc *GoodController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := ExtractToken(c)                  // id pemilik dari token
		goodID, err := strconv.Atoi(c.Param("id")) // id barang dari uri parameter
		if err != nil {
			log.Println("convert id error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gunakan input angka",
			})
		}

		err = gc.Mdl.Delete(userID, goodID)
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sukses menghapus data",
		})
	}
}
