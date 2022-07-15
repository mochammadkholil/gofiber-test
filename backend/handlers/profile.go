package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"backendsvc/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type ProfileHandler struct {
	DB *sql.DB
}

func NewProfileHandler(db *sql.DB) *ProfileHandler {
	return &ProfileHandler{
		DB: db,
	}
}

func (r *ProfileHandler) Login(ctx *fiber.Ctx) error {
	response := models.ApiResponse{}
	login := models.Login{}
	profile := models.Profile{}
	errorMsg := ""
	if err := ctx.BodyParser(&login); err != nil {
		errorMsg = fmt.Sprintf("An error occured: %v", err)
	} else {
		if login.Email != "" {
			row := r.DB.QueryRow("SELECT nama, email, no_telp FROM users WHERE email=$1 AND password=md5($2)", login.Email, login.Password)
			if row.Err() != nil {
				log.Printf("An error occured while executing query: %v", err)
				errorMsg = "internal server error"
			} else {
				e := row.Scan(&profile.Nama, &profile.Email, &profile.NoTelp)
				if e != nil {
					errorMsg = fmt.Sprintf("login failed!, %v", e)
				} else {
					// Create the Claims
					claims := jwt.MapClaims{
						"name":  profile.Nama,
						"email": profile.Email,
						"admin": true,
						"exp":   time.Now().Add(time.Hour * 72).Unix(),
					}

					// Create token
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
					// Generate encoded token and send it as response.
					t, err := token.SignedString([]byte("secret"))
					if err != nil {
						return ctx.SendStatus(fiber.StatusInternalServerError)
					}
					response.Success = true
					response.Message = "login success"
					response.Data = t
				}
			}
		} else {
			errorMsg = "invalid email or password"
		}
	}

	if errorMsg != "" {
		response.Message = errorMsg
	}
	return ctx.JSON(response)
}

func (r *ProfileHandler) Register(ctx *fiber.Ctx) error {
	response := models.ApiResponse{}
	var errorMsg string
	profile := models.Profile{}

	if err := ctx.BodyParser(&profile); err != nil {
		errorMsg = fmt.Sprintf("An error occured: %v", err)
	} else {

		if profile.Email != "" && profile.Nama != "" && profile.NoTelp != "" && profile.Password != "" {
			_, err := r.DB.Exec("INSERT INTO users(nama, email, no_telp, password) VALUES($1, $2, $3, md5($4))",
				profile.Nama, profile.Email, profile.NoTelp, profile.Password)

			if err != nil {
				log.Printf("An error occured while executing query: %v", err)
				errorMsg = "internal server error"
			} else {
				errorMsg = ""
			}
		} else {
			errorMsg = "invalid payload"
		}
	}

	if errorMsg == "" {
		profile.Password = "*****"
		response.Success = true
		response.Data = profile
		response.Message = "registration success"
	} else {
		response.Message = errorMsg
	}

	return ctx.JSON(response)
}

func (r *ProfileHandler) GetProfile(ctx *fiber.Ctx) error {
	response := models.ApiResponse{}
	var errorMsg string
	email := ctx.Query("email")
	profile := models.Profile{}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email_1 := claims["email"].(string)

	if email_1 != email {
		errorMsg = "don't have permission"
	} else {
		row := r.DB.QueryRow("SELECT nama, email, no_telp FROM users WHERE email=$1", email)
		if row.Err() != nil {
			log.Printf("An error occured while executing query: %v", row.Err())
			errorMsg = "internal server error"
		} else {
			e := row.Scan(&profile.Nama, &profile.Email, &profile.NoTelp)
			if e != nil {
				errorMsg = fmt.Sprintf("fail get profile data, %v", e.Error())
			}
		}
	}
	if errorMsg != "" {
		response.Message = errorMsg
	} else {
		profile.Password = "*****"
		response.Success = true
		response.Data = profile
		response.Message = "get profile data success"
	}

	return ctx.JSON(response)
}

func (r *ProfileHandler) UpdateProfile(ctx *fiber.Ctx) error {
	response := models.ApiResponse{}
	profile := models.Profile{}
	var errorMsg string
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email_1 := claims["email"].(string)

	if email_1 == "" {
		errorMsg = "don't have permission"
	} else {
		if err := ctx.BodyParser(&profile); err != nil {
			errorMsg = fmt.Sprintf("An error occured: %v", err)
		} else {
			var err error
			if profile.Password == "" {
				_, err = r.DB.Exec("UPDATE users SET nama=$1, email=$2, no_telp=$3 WHERE email=$2",
					profile.Nama, profile.Email, profile.NoTelp)
			} else {
				_, err = r.DB.Exec("UPDATE users SET nama=$1, email=$2, no_telp=$3, password=$4 WHERE email=$2",
					profile.Nama, profile.Email, profile.NoTelp, profile.Password)
			}

			if err != nil {
				log.Printf("An error occured while executing query: %v", err)
				errorMsg = "internal server error"
			}
		}
	}

	if errorMsg != "" {
		response.Message = errorMsg
	} else {
		profile.Password = "*****"
		response.Success = true
		response.Data = profile
		response.Message = "update profile data success"
	}

	return ctx.JSON(response)
}
