package models

import "gorm.io/gorm"

type tOrders struct {
	ID      uint    `gorm:"primary key;autoIncrement" json:"id"`
	OrderId *string `gorm:"primary key" json:"NomorLab"`
	PatId   *string `json:"Nomor Rekam Medis"`
	PatName *string `json:"NamaPasien"`
	Address *string `json:"Alamat"`
}

type tResults struct {
	ID       uint     `gorm:"primary key;autoIncrement" json:"id"`
	rOrderId *tOrders `json:"Nomorlab"`
}

func MigratetOrders(db *gorm.DB) error {
	err := db.AutoMigrate(&tOrders{})
	return err
}

func MigratetResults(db *gorm.DB) error {
	err := db.AutoMigrate(&tResults{})
	return err
}
