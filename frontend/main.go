package main

import (
	"fmt"
	"frontend/models"
	"frontend/services"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var Token string

func main() {
	//tmpl, err := template.ParseFiles("layout.html", "_navigasi.html")

	engine := html.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "1111"
	}

	app.Static("/", "./public") // add this before starting the app
	app.Get("/login", func(c *fiber.Ctx) error { return c.Render("login", nil) })
	app.Get("/register", func(c *fiber.Ctx) error { return c.Render("register", nil) })
	app.Get("/produk", func(c *fiber.Ctx) error { return c.Render("produk", nil) })

	app.Post("/register", func(c *fiber.Ctx) error {
		jsondata := models.Profile{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Nama:     c.FormValue("nama"),
			NoTelp:   c.FormValue("no_telp"),
		}

		service := services.NewServices()
		resp := service.Register(jsondata)
		if resp.Success {
			return c.Redirect("/login")
		}

		return c.JSON(resp)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		jsondata := models.Login{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		service := services.NewServices()
		resp := service.Login(jsondata)
		if resp.Success {
			//Token = resp.Data.(string)
			cookie := fiber.Cookie{
				Name:     "auth",
				Value:    resp.Data.(string),
				Expires:  time.Now().Add(time.Hour * 24),
				HTTPOnly: true,
			}

			c.Cookie(&cookie)

			return c.Redirect("/")
		} else {
			return c.SendString("don't have permission")
		}
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		cookie := fiber.Cookie{
			Name:     "auth",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		return c.Redirect("/login")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		e := Auth(c)
		if e != nil {
			return c.Redirect("/login")
		} else {
			return c.Next()
		}
	}, func(c *fiber.Ctx) error {
		token := c.Cookies("auth")
		service := services.NewServices()
		resp := service.ListAllProduk(token)
		data, ok := resp.Data.([]interface{})
		barang := []map[string]string{}
		if ok {
			for i := 0; i < len(data); i++ {
				item := data[i].(map[string]interface{})
				barang = append(barang, map[string]string{
					"nama":   item["nama"].(string),
					"sku":    item["sku"].(string),
					"jumlah": fmt.Sprintf("%v", item["jumlah"].(float64)),
				})
			}
		}

		kue := c.Cookies("username")
		return c.Render("index", fiber.Map{
			"username": kue,
			"barang":   barang,
		})
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		token := c.Cookies("auth")
		email := c.Cookies("email")
		service := services.NewServices()
		resp := service.GetProfile(email, token)
		if resp.Success {
			data := resp.Data.(map[string]interface{})

			return c.Render("profile", fiber.Map{
				"nama":    data["nama"].(string),
				"no_telp": data["no_telp"].(string),
				"email":   data["email"].(string),
			})
		} else {
			return c.JSON(resp)
		}
	})

	app.Post("/profile", func(c *fiber.Ctx) error {
		token := c.Cookies("auth")
		jsondata := models.Profile{
			Email:  c.FormValue("email"),
			Nama:   c.FormValue("nama"),
			NoTelp: c.FormValue("no_telp"),
		}

		service := services.NewServices()
		resp := service.Update(jsondata, token)
		if resp.Success {
			return c.Redirect("/")
		}

		return c.JSON(resp)
	})

	app.Post("/produk", func(c *fiber.Ctx) error {
		token := c.Cookies("auth")
		jumlah, e := strconv.ParseInt(c.FormValue("jumlah"), 10, 32)
		if e != nil {
			return c.SendString("invalid number")
		}

		jsondata := models.Produk{
			Nama:   c.FormValue("nama"),
			SKU:    c.FormValue("sku"),
			Jumlah: jumlah,
		}

		service := services.NewServices()
		resp := service.Produk(jsondata, token)
		if resp.Success {
			return c.Redirect("/")
		}

		return c.JSON(resp)
	})

	log.Fatalln(app.Listen(fmt.Sprintf("localhost:%v", port)))
}

func Auth(c *fiber.Ctx) error {
	cookie := c.Cookies("auth")
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return fmt.Errorf("unauthenticated")
	} else {
		claims := token.Claims.(jwt.MapClaims)
		kue := fiber.Cookie{
			Name:     "username",
			Value:    claims["name"].(string),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}
		c.Cookie(&kue)

		kue = fiber.Cookie{
			Name:     "email",
			Value:    claims["email"].(string),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}
		c.Cookie(&kue)

		c.Status(fiber.StatusOK)
	}

	return nil
}
