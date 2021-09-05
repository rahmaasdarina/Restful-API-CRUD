package repository

import (
	"restful-crud-api/model"

	"gorm.io/gorm"
)

type ProdukRepository struct {
	DB *gorm.DB
}

func (r *ProdukRepository) FindAll() ([]model.Produk, error) {
	var produks []model.Produk
	findResult := r.DB.Find(&produks)

	return produks, findResult.Error
}

func (r *ProdukRepository) Save(produks model.Produk) (model.Produk, error) {
	//Query Transaction
	errOuter := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Omit("ID").Create(&produks).Error

		if err != nil {
			return err
		}

		err = tx.Omit("ID").Create(&produks).Error //Omit di ga akan kirim ID

		if err != nil {
			return err
		}

		return nil
	})
	return produks, errOuter

	//Query tanpa transaction
	// trx := r.DB.Create(&produks)
	// return produks, trx.Error
}

func (r *ProdukRepository) FindById(id string) (model.Produk, error) {
	produks := &model.Produk{}
	findResult := r.DB.First(&produks, id)

	return *produks, findResult.Error
}

func (r *ProdukRepository) DeleteById(id string) error {
	produks := &model.Produk{}
	result := r.DB.Delete(&produks, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ProdukRepository) UpdateById(produks *model.Produk, id string) error {

	result := r.DB.Model(&produks).Where("id = ?", id).Updates(map[string]interface{}{"name": produks.Name, "qty": produks.Qty, "supplier": produks.Supplier})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
