package api

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *API) DeleteProductById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete products by Id DELETE /api/v1/products/{id}")
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

	err = api.storage.Product().DeleteProducts(id)
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
		Message:    fmt.Sprintf("products with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) PostProduct(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	var brand models.Brand

	prod := models.Product{
		Brand: &brand,
	}

	api.logger.Info("Post User POST /api/v1/products")
	err := json.NewDecoder(req.Body).Decode(&prod)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		fmt.Println()
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(prod.Brand.Id)
	fmt.Println(prod.Slug)
	fmt.Println(prod.Name)
	fmt.Println(prod.Full_description)
	fmt.Println(prod.Short_description)
	fmt.Println(prod.SKU)
	err = api.storage.Product().CreateProduct(&prod)
	if err != nil {
		api.logger.Info("Troubles while creating new product:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: 200,
		Message:    "Created",
		IsError:    true,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) GetProduct(writer http.ResponseWriter, req *http.Request) {
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
		api.logger.Info("Invalid json recieved from products")
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
	brand, err := api.storage.Product().FilterAllProducts(&filter)
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
	Product := struct {
		Items  []*models.Product
		Filter models.Filter
	}{
		Items:  brand,
		Filter: filter,
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(Product)
}
func (api *API) PutProducts(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	var brand models.Brand
	prod := models.Product{
		Brand: &brand,
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
	err = json.NewDecoder(req.Body).Decode(&prod)
	fmt.Println(prod)
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
	err = api.storage.Product().UpdateProduct(id, &prod)
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
func (api *API) GetProductById(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("start GetProduct")

	var product *models.Product

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
	product, err = api.storage.Product().GetProductById(id)
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
	json.NewEncoder(writer).Encode(product)
}
