package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (api *API) DeleteProducts_videosById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Products_videos By Id /api/v1/products_videos")

	var products_videos models.Products_videos
	log.Println(products_videos)

	err := json.NewDecoder(req.Body).Decode(&products_videos)
	if err != nil {
		api.logger.Info("Invalid json recieved from Products_videos")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(products_videos)
	err = api.storage.Products_videos().DeleteProduct_VideoById(&products_videos)
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
	} else {
		msg := Message{
			StatusCode: 200,
			Message:    "delete complited",
			IsError:    false,
		}
		writer.WriteHeader(201)
		json.NewEncoder(writer).Encode(msg)
	}
}
func (api *API) PostProduct_video(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	brand := models.Brand{}
	product := models.Product{
		Brand: &brand,
	}
	var (
		video models.Videos
	)
	prod := models.Products_videos{
		Product: &product,
		Videos:  &video,
	}

	api.logger.Info("Post Products_videos POST /api/v1/products_videos")
	err := json.NewDecoder(req.Body).Decode(&prod)
	if err != nil {
		api.logger.Info("Invalid json recieved from products_videos")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.Products_videos().CreateProduct_video(&prod)
	if err != nil {
		api.logger.Info("Troubles while creating new products_videos:", err)
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
		Message:    "Products_videos Created",
		IsError:    true,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) GetProducts_videos(writer http.ResponseWriter, req *http.Request) {
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
	brand, err := api.storage.Products_videos().FilterAllProducts_video(&filter)
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
