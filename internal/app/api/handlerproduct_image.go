package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (api *API) DeleteProduct_imageById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete products_images By Id /api/v1/products_images")

	var products_images models.Products_images
	log.Println(products_images)

	err := json.NewDecoder(req.Body).Decode(&products_images)
	if err != nil {
		api.logger.Info("Invalid json recieved from products_images")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(products_images)
	err = api.storage.Products_images().DeleteProduct_imageById(&products_images)
	if err != nil {
		api.logger.Info("Troubles while connections to the warehouse database:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 200,
		Message:    "delete complited",
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
func (api *API) PostProduct_image(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	brand := models.Brand{}
	product := models.Product{
		Brand: &brand,
	}

	var (
		image models.Images
	)
	prod := models.Products_images{
		Product: &product,
		Image:   &image,
	}

	api.logger.Info("Post Products_images POST /api/v1/products_images")
	err := json.NewDecoder(req.Body).Decode(&prod)
	if err != nil {
		api.logger.Info("Invalid json recieved from products_images")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.Products_images().CreateProduct_image(&prod)
	if err != nil {
		api.logger.Info("Troubles while creating new products_images:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}

	msg := Message{
		StatusCode: 200,
		Message:    "Products_images created",
		IsError:    true,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) GetProducts_images(writer http.ResponseWriter, req *http.Request) {
	var (
		filter models.Filter
	)
	initHeaders(writer)
	pg := models.Pages{}
	fl := make([]models.FieldFilter, 0)
	so := make([]models.FieldSort, 0)
	filter = models.Filter{
		Fields: &fl,
		Sorts:  &so,
		Pages:  &pg,
	}
	err := json.NewDecoder(req.Body).Decode(&filter)
	if err != nil {
		api.logger.Info("Invalid json recieved from brands")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(filter)
	brand, err := api.storage.Products_images().FilterAllProducts_images(&filter)
	if err != nil {
		api.logger.Info("Error while brands SelectAll: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(brand)
	json.NewEncoder(writer).Encode(filter)
}
