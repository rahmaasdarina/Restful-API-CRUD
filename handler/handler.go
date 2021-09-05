package handler

import (
	"restful-crud-api/model"
	"restful-crud-api/service"

	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type ProdukHandler struct {
	ProdukService service.ProdukService
}

func (r *ProdukHandler) GetProduks(c *fiber.Ctx) {
	produks, err := r.ProdukService.GetAllProduk()

	if err != nil {
		c.Status(403).Send(err)
		return
	}
	c.JSON(produks)
}

func (r *ProdukHandler) InsertProduk(c *fiber.Ctx) {
	produk := &model.Produk{}

	if err := c.BodyParser(produk); err != nil {
		c.Status(403).Send(err)
		return
	}
	r.ProdukService.InsertProduk(*produk)
	c.JSON(produk)
}

func (r *ProdukHandler) GetProdukById(c *fiber.Ctx) {
	id := c.Params("id")
	produks, err := r.ProdukService.GetProdukById(id)

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "data not found",
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "success retrieve data",
		"result": produks,
	})
}

func (r *ProdukHandler) DeleteProdukById(c *fiber.Ctx) {
	id := c.Params("id")
	err := r.ProdukService.DeleteById(id)

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "data not found",
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "success delete data",
	})

}

func (r *ProdukHandler) UpdateProdukById(c *fiber.Ctx) {
	produk := &model.Produk{}
	id := c.Params("id")
	var mysqlErr *mysql.MySQLError

	if err := c.BodyParser(produk); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": 400,
			"error":  true,
			"msg":    err.Error(),
		})
		return
	}

	produk_updt, err := r.ProdukService.UpdateProduct(produk, id)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"status": 422,
				"error":  true,
				"msg":    err.Error(),
			})
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": 404,
				"error":  true,
				"msg":    "data not found",
			})
			return
		}

		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": 500,
			"error":  true,
			"msg":    err.Error(),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"error":  false,
		"msg":    "success update data",
		"result": produk_updt,
	})

}
