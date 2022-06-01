package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *API) DeleteImageById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Image by Id DELETE /api/v1/image/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.Image().DeleteImage(id)
	if err != nil {
		api.logger.Info("Troubles while deleting database elemnt from table with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Image with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) PostImages(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post image POST /api/v1/images")
	var image models.Images
	err := json.NewDecoder(req.Body).Decode(&image)
	if err != nil {
		api.logger.Info("Invalid json recieved from image")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Image().CreateImage(&image)
	if err != nil {
		api.logger.Info("Troubles while creating new image:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}
func (api *API) PutImages(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	var image models.Images
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Unable to parse request target")
		msg := Message{
			StatusCode: 400,
			Message:    "Bad request id",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(req.Body).Decode(&image)
	fmt.Println(image)
	if err != nil {
		api.logger.Info("Invalid request body json")
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid request body json",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.Image().Update(id, &image)
	if err != nil {
		api.logger.Info("Failed to update target", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Updating failed",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 200,
		Message:    "Update successfull",
		IsError:    false,
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) GetImages(writer http.ResponseWriter, req *http.Request) {
	var (
		filter models.PageRequest
	)
	initHeaders(writer)
	fl := make([]models.Field, 0)

	filter = models.PageRequest{
		Fields: &fl,
	}
	fmt.Println(req.Body)
	err := json.NewDecoder(req.Body).Decode(&filter)
	if err != nil {
		api.logger.Info("Invalid json recieved from images")
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
	list, err := api.storage.Image().FilterAllImages(&filter)
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
	Resp := struct {
		PgNum    int `json:"pg_number"`
		PgLen    int `json:"pg_length"`
		TotalRec int `json:"total_rec"`
		TotalPg  int `json:"total_pg"`
		List     []*models.Images
	}{
		filter.PageNumber,
		filter.PageLength,
		filter.TotalRecords,
		AllPage(filter.TotalRecords, filter.PageLength),
		list,
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(Resp)
}
