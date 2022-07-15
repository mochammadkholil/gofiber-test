package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"backendsvc/models"

	"github.com/gofiber/fiber/v2"
)

type ProdukHandler struct {
	DB *sql.DB
}

func NewProdukHandler(db *sql.DB) *ProdukHandler {
	return &ProdukHandler{
		DB: db,
	}
}

func (r *ProdukHandler) GetAllProduk(ctx *fiber.Ctx) error {
	response := models.ApiResponse{}
	var errorMsg string
	var produks []models.Produk

	rows, err := r.DB.Query("SELECT id, nama, sku, jumlah FROM produk")
	if err != nil {
		log.Printf("An error occured while executing query: %v", err)
		errorMsg = "internal server error"
	} else {
		defer rows.Close()

		for rows.Next() {
			var p models.Produk
			err := rows.Scan(&p.Id, &p.Nama, &p.SKU, &p.Jumlah)
			if err != nil {
				log.Printf("An error occured while scanning data: %v", err)
				errorMsg = "internal server error"
			}
			produks = append(produks, p)
		}
	}

	if errorMsg != "" {
		response.Message = errorMsg
	} else {
		response.Success = true
		response.Data = produks
		response.Message = "fetch produk data success"
	}

	return ctx.JSON(response)
}

func (r *ProdukHandler) InputProduk(ctx *fiber.Ctx) error {
	response := models.ApiResponse{}
	var errorMsg string
	produk := models.Produk{}

	if err := ctx.BodyParser(&produk); err != nil {
		errorMsg = fmt.Sprintf("An error occured: %v", err)
	} else {
		if produk.Nama != "" && produk.SKU != "" {
			_, err := r.DB.Exec("INSERT INTO produk(nama, sku, jumlah) VALUES ($1, $2, $3)",
				produk.Nama, produk.SKU, produk.Jumlah)

			if err != nil {
				log.Printf("An error occured while executing query: %v", err)
				errorMsg = "internal server error"
			}
		} else {
			errorMsg = "invalid payload"
		}
	}

	if errorMsg != "" {
		response.Message = errorMsg
	} else {
		response.Success = true
		response.Data = produk
		response.Message = "input produk data success"
	}

	return ctx.JSON(response)
}

func (r *ProdukHandler) UpdateProduk(ctx *fiber.Ctx) error {
	response := models.ApiResponse{}
	var errorMsg string

	p := models.Produk{}
	if err := ctx.BodyParser(&p); err != nil {
		errorMsg = fmt.Sprintf("An error occured: %v", err)
	} else {
		_, err := r.DB.Exec("UPDATE produk SET nama=$1, sku=$2, jumlah=$3 WHERE id=$4",
			p.Nama, p.SKU, p.Jumlah, p.Id)

		if err != nil {
			log.Printf("An error occured while executing query: %v", err)
			errorMsg = "internal server error"
		}
	}

	if errorMsg != "" {
		response.Message = errorMsg
	} else {
		response.Success = true
		response.Data = p
		response.Message = "update produk data success"
	}

	return ctx.JSON(response)
}
