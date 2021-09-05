package service

import (
	"restful-crud-api/model"
	"restful-crud-api/repository"
)

type ProdukService struct {
	ProdukRepo repository.ProdukRepository
}

func (s *ProdukService) GetAllProduk() ([]model.Produk, error) {
	produks, err := s.ProdukRepo.FindAll()

	return produks, err
}

func (s *ProdukService) InsertProduk(produks model.Produk) error {
	_, err := s.ProdukRepo.Save(produks)

	return err
}

func (s *ProdukService) GetProdukById(id string) (model.Produk, error) {
	produks, err := s.ProdukRepo.FindById(id)
	if err != nil {
		return produks, err
	}

	return produks, nil
}

func (s *ProdukService) DeleteById(id string) error {
	err := s.ProdukRepo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProdukService) UpdateProduct(produks *model.Produk, id string) (product model.Produk, err error) {
	err = s.ProdukRepo.UpdateById(produks, id)
	if err != nil {
		return product, err
	}

	// select movie
	product, err = s.ProdukRepo.FindById(id)
	if err != nil {
		return product, err
	}

	return product, nil
}
