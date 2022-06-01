package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *API) PostBrand(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post brands POST /api/v1/brands")
	var brand models.Brand
	err := json.NewDecoder(req.Body).Decode(&brand)
	if err != nil {
		api.logger.Info("Invalid json recieved from brand")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Brand().CreateBrand(&brand)
	if err != nil {
		api.logger.Info("Troubles while creating new brand:", err)
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
func (api *API) DeleteBrandById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete brand by Id DELETE /api/v1/brands/{id}")
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
	err = api.storage.Brand().DeleteBrand(id)
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
		Message:    fmt.Sprintf("brands with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) PutBrand(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	var brand models.Brand
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
	err = json.NewDecoder(req.Body).Decode(&brand)
	fmt.Println(brand)
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
	err = api.storage.Brand().UpdateBrand(id, &brand)
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
func (api *API) GetBrands(writer http.ResponseWriter, req *http.Request) {

	fmt.Println("start GetBrands")
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
	list, err := api.storage.Brand().FilterAllBrands(&filter)
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
		List     []*models.Brand
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
func (api *API) GetBrandsById(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("start GetBrands")
	var (
		brand *models.Brand
	)
	initHeaders(writer)
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
	brand, err = api.storage.Brand().GetBrandById(id)
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
}
