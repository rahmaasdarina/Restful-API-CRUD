package main

import (
	"restful-crud-api/handler"
	"restful-crud-api/model"
	"restful-crud-api/repository"
	"restful-crud-api/service"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	//app.Get("/get-produk", getProduk)

	DbConn, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/produk_api?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic(err.Error())
	}

	model.MigrateProduk(DbConn)

	produkRepo := repository.ProdukRepository{
		DB: DbConn,
	}

	produkService := service.ProdukService{
		ProdukRepo: produkRepo,
	}

	produkHandler := handler.ProdukHandler{
		ProdukService: produkService,
	}

	app.Get("/get-produk", produkHandler.GetProduks)
	app.Post("/create-produk", produkHandler.InsertProduk)
	app.Get("/get-produk/:id", produkHandler.GetProdukById)
	app.Delete("/delete-produk/:id", produkHandler.DeleteProdukById)
	app.Put("/update-produk/:id", produkHandler.UpdateProdukById)

	app.Listen(3000)
}
