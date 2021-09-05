package model

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	Name     string `json:"name"`
	Qty      int    `json:"qty"`
	Supplier string `json:"supplier"`
}

func MigrateProduk(db *gorm.DB) {
	db.AutoMigrate(&Produk{})
}
