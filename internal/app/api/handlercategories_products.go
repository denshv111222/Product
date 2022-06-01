package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (api *API) DeleteCategories_productsById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Categories_products /api/v1/categories_products")

	var cat_prod models.Categories_products
	log.Println(cat_prod)

	err := json.NewDecoder(req.Body).Decode(&cat_prod)
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
	fmt.Println(cat_prod)
	err = api.storage.Categories_products().DeleteCategories_productsById(&cat_prod)
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
func (api *API) PostCategories_products(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	brand := models.Brand{}
	product := models.Product{
		Brand: &brand,
	}
	categ := models.Categories{}
	cat_prod := models.Categories_products{
		Product:    &product,
		Categories: &categ,
	}

	api.logger.Info("Post Categories_products POST /api/v1/categories_products")
	err := json.NewDecoder(req.Body).Decode(&cat_prod)
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
	err = api.storage.Categories_products().CreateCategories_products(&cat_prod)
	if err != nil {
		api.logger.Info("Troubles while creating new category_product:", err)
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
		Message:    "Category_product Created",
		IsError:    true,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) GetCategories_products(writer http.ResponseWriter, req *http.Request) {
	var (
		filter models.PageRequest
	)
	initHeaders(writer)
	fl := make([]models.Field, 0)

	filter = models.PageRequest{
		Fields: &fl,
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
	list, err := api.storage.Categories_products().FilterAllCategories_products(&filter)
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
		List     []*models.Categories_products
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
