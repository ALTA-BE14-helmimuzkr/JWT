package model

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Good struct {
	gorm.Model
	UserID int   `json:"user_id"`
	User   *User `json:"-" gorm:"constraint"`
	// User User   ` gorm:"constraint"`
	Name string `json:"name" form:"nama"`
	Qty  int    `json:"qty" form:"qty"`
}

type GoodModel struct {
	DB *gorm.DB
}

// Menambahkan barang baru
func (gm *GoodModel) Insert(good Good) (Good, error) {
	tx := gm.DB.Create(&good)
	if tx.Error != nil {
		return Good{}, tx.Error
	}

	return good, nil
}

// Mendapatkan barang yang dipunyai oleh pemilik
func (gm *GoodModel) GetAll(userID int) ([]Good, error) {
	goods := []Good{}
	tx := gm.DB.
		Where("goods.user_id = ?", userID).
		Joins("JOIN users ON users.id = goods.id").
		Find(&goods)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return goods, nil
}

// Mendapatkan barang punya pemilik yang dipilih menggunakan id barang
func (gm *GoodModel) GetByID(goodID int, userID int) (Good, error) {
	good := Good{}
	tx := gm.DB.
		Where("goods.user_id = ? AND goods.id = ?", userID, goodID).
		Joins("JOIN users ON users.id = goods.id").
		Find(&good)
	if tx.Error != nil {
		return Good{}, tx.Error
	}
	return good, nil
}

// Mengubah barang, dimana barang tersebut adalah punya pemilik
func (gm *GoodModel) Update(good Good) (Good, error) {
	tx := gm.DB.
		Where("user_id = ? AND id = ?", good.UserID, good.ID).
		Updates(&good)
	if tx.Error != nil {
		return Good{}, tx.Error
	}
	return good, nil
}

// Menghapus barang milik pemilik
func (gm *GoodModel) Delete(goodID int, userID int) error {
	qry := gm.DB.
		Where("user_id = ? AND id = ?", userID, goodID).
		Delete(&Good{})

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("no data processed")
		return errors.New("tidak ada data yang dihapus")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("tidak bisa menghapus data")
	}

	return nil
}
