package api

import (
	"backendsvc/extensions"
	"backendsvc/handlers"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, db *sql.DB) *fiber.App {
	v1 := app.Group("/api/v1")
	{
		v1.Use(extensions.Auth([]string{"login", "register"}))

		v1.Post("/login", handlers.NewProfileHandler(db).Login)
		v1.Post("/register", handlers.NewProfileHandler(db).Register)

		profile := v1.Group("/profile")
		{
			profile.Get("", handlers.NewProfileHandler(db).GetProfile)
			profile.Put("", handlers.NewProfileHandler(db).UpdateProfile)
		}

		produk := v1.Group("/produk")
		{
			produk.Get("", handlers.NewProdukHandler(db).GetAllProduk)
			produk.Post("", handlers.NewProdukHandler(db).InputProduk)
			produk.Put("", handlers.NewProdukHandler(db).UpdateProduk)
		}
	}

	return app
}
