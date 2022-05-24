package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (api *API) DeleteAttributes_values_productsById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Attributes_values_products Att /api/v1/attributes_values_products")

	var attributes_values_products models.Attributes_values_products
	log.Println(attributes_values_products)

	err := json.NewDecoder(req.Body).Decode(&attributes_values_products)
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
	fmt.Println(attributes_values_products)
	err = api.storage.Attributes_values_products().DeleteAttributes_values_productsById(&attributes_values_products)
	if err != nil {
		api.logger.Info("Troubles while connections to the Products database:", err)
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
func (api *API) PostAttributes_values_products(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	unit := models.Units{}
	attr := models.Attributes{
		Units: &unit,
	}
	attr_val := models.Attributes_values{
		Attributes: &attr,
	}
	brand := models.Brand{}
	product := models.Product{
		Brand: &brand,
	}
	attr_val_prod := models.Attributes_values_products{
		Produkt:           &product,
		Attributes_values: &attr_val,
	}

	api.logger.Info("Post Attributes_values_products POST /api/v1/attributes_values_products")
	err := json.NewDecoder(req.Body).Decode(&attr_val_prod)
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
	err = api.storage.Attributes_values_products().CreateAttributes_values_products(&attr_val_prod)
	if err != nil {
		api.logger.Info("Troubles while creating new attribute_value_product:", err)
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
		Message:    "attributes_values_products Created",
		IsError:    true,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) GetAettributes_values_products(writer http.ResponseWriter, req *http.Request) {
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
	brand, err := api.storage.Attributes_values_products().FilterAllAttributes_values_products(&filter)
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
