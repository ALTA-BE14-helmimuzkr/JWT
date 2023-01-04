package model

import "gorm.io/gorm"

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

func (gm *GoodModel) Insert(good Good) (Good, error) {
	tx := gm.DB.Create(&good)
	if tx.Error != nil {
		return Good{}, tx.Error
	}

	return good, nil
}

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

// func (gm *GoodModel) GetByID(goodID int, userID int) (Good, error) {
// 	good := Good{}
// 	tx := gm.DB.Where()
// }
