package services

import (
	"encoding/json"
	"fmt"
	"frontend/models"
	"log"

	"github.com/go-resty/resty/v2"
)

type Services struct {
}

func NewServices() *Services {
	return &Services{}
}

func (s Services) Login(data models.Login) *models.Response {
	url := "http://localhost:2222/api/v1/login"
	client := resty.New()
	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)

	resp := models.Response{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		errMsg := fmt.Sprintf("error while login, %v", err)
		log.Print(errMsg)
		resp.Message = errMsg
		return &resp
	}

	return &resp
}

func (s Services) Register(data models.Profile) *models.Response {
	url := "http://localhost:2222/api/v1/register"
	client := resty.New()
	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)

	resp := models.Response{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		errMsg := fmt.Sprintf("error while login, %v", err)
		log.Print(errMsg)
		resp.Message = errMsg
		return &resp
	}

	return &resp
}

func (s Services) GetProfile(email string, token string) *models.Response {
	url := "http://localhost:2222/api/v1/profile"
	client := resty.New()
	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("email", email).
		SetHeader("Authorization", "Bearer "+token).
		Get(url)

	resp := models.Response{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		errMsg := fmt.Sprintf("error while login, %v", err)
		log.Print(errMsg)
		resp.Message = errMsg
		return &resp
	}

	return &resp
}

func (s Services) Update(data models.Profile, token string) *models.Response {
	url := "http://localhost:2222/api/v1/profile"
	client := resty.New()
	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		SetBody(data).
		Put(url)

	resp := models.Response{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		errMsg := fmt.Sprintf("error while login, %v", err)
		log.Print(errMsg)
		resp.Message = errMsg
		return &resp
	}

	return &resp
}

func (s Services) ListAllProduk(token string) *models.Response {
	url := "http://localhost:2222/api/v1/produk"
	client := resty.New()
	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		Get(url)

	resp := models.Response{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		errMsg := fmt.Sprintf("error while login, %v", err)
		log.Print(errMsg)
		resp.Message = errMsg
		return &resp
	}

	return &resp
}

func (s Services) Produk(data models.Produk, token string) *models.Response {
	url := "http://localhost:2222/api/v1/produk"
	client := resty.New()
	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		SetBody(data).
		Post(url)

	resp := models.Response{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		errMsg := fmt.Sprintf("error while login, %v", err)
		log.Print(errMsg)
		resp.Message = errMsg
		return &resp
	}

	return &resp
}
