package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *API) DeleteAttributeById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete attribute by Id DELETE /api/v1/attributes/{id}")
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
	err = api.storage.Attributes().DeleteAttributes(id)
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
		Message:    fmt.Sprintf("Attributes with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) PostAttributes(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	var unit models.Units

	attr := models.Attributes{
		Units: &unit,
	}

	api.logger.Info("Post Attributes POST /api/v1/attributes")
	err := json.NewDecoder(req.Body).Decode(&attr)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.Attributes().CreateAttributes(&attr)
	if err != nil {
		api.logger.Info("Troubles while creating new attributes:", err)
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
		Message:    "Attributes Created",
		IsError:    true,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) PutAttribute(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	var unit models.Units

	attr := models.Attributes{
		Units: &unit,
	}
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
	err = json.NewDecoder(req.Body).Decode(&attr)
	fmt.Println(attr)
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
	err = api.storage.Attributes().UpdateAttribute(id, &attr)
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
func (api *API) GetAttributes(writer http.ResponseWriter, req *http.Request) {
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
	brand, err := api.storage.Attributes().FilterAllAttributes(&filter)
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
