package api

import (
	"GitHab/Standart_Server_API/storage"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix = "/api/v2"
)

func (a *API) configreLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configreRouterField() {
	a.router.HandleFunc(prefix+"/images", a.GetImages).Methods("GET")
	a.router.HandleFunc(prefix+"/images", a.PostImages).Methods("POST")
	a.router.HandleFunc(prefix+"/images/{id}", a.PutImages).Methods("PUT")
	a.router.HandleFunc(prefix+"/images/{id}", a.DeleteImageById).Methods("DELETE")

	a.router.HandleFunc(prefix+"/videos", a.GetVideos).Methods("GET")
	a.router.HandleFunc(prefix+"/videos", a.PostVideo).Methods("POST")
	a.router.HandleFunc(prefix+"/videos/{id}", a.PutVideo).Methods("PUT")
	a.router.HandleFunc(prefix+"/videos/{id}", a.DeleteVideoById).Methods("DELETE")

	a.router.HandleFunc(prefix+"/units", a.GetUnits).Methods("GET")
	a.router.HandleFunc(prefix+"/units", a.PostUnit).Methods("POST")
	a.router.HandleFunc(prefix+"/units/{id}", a.PutImages).Methods("PUT")
	a.router.HandleFunc(prefix+"/units/{id}", a.DeleteUnitById).Methods("DELETE")

	a.router.HandleFunc(prefix+"/brands", a.PostBrand).Methods("POST")
	a.router.HandleFunc(prefix+"/brands/{id}", a.PutBrand).Methods("PUT")
	a.router.HandleFunc(prefix+"/brands/{id}", a.DeleteBrandById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/brands", a.GetBrands).Methods("GET")
	a.router.HandleFunc(prefix+"/brand/{id}", a.GetBrandsById).Methods("GET")

	a.router.HandleFunc(prefix+"/categories", a.GetCategories).Methods("GET")
	a.router.HandleFunc(prefix+"/category/{id}", a.GetCategoryById).Methods("GET")
	a.router.HandleFunc(prefix+"/categories", a.PostCategory).Methods("POST")
	a.router.HandleFunc(prefix+"/categories/{id}", a.PutCategory).Methods("PUT")
	a.router.HandleFunc(prefix+"/categories/{id}", a.DeleteCategoryById).Methods("DELETE")

	a.router.HandleFunc(prefix+"/products", a.GetProduct).Methods("GET")
	a.router.HandleFunc(prefix+"/product/{id}", a.GetProductById).Methods("GET")
	a.router.HandleFunc(prefix+"/products/{id}", a.DeleteProductById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/products/{id}", a.PutProducts).Methods("PUT")
	a.router.HandleFunc(prefix+"/products", a.PostProduct).Methods("POST")

	a.router.HandleFunc(prefix+"/products_images", a.GetProducts_images).Methods("GET")
	a.router.HandleFunc(prefix+"/products_images", a.DeleteProduct_imageById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/products_images", a.PostProduct_image).Methods("POST")

	a.router.HandleFunc(prefix+"/products_videos", a.GetProducts_videos).Methods("GET")
	a.router.HandleFunc(prefix+"/products_videos", a.DeleteProducts_videosById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/products_videos", a.PostProduct_video).Methods("POST")

	a.router.HandleFunc(prefix+"/attributes", a.GetAttributes).Methods("GET")
	a.router.HandleFunc(prefix+"/attributes/{id}", a.DeleteAttributeById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/attributes", a.PostAttributes).Methods("POST")
	a.router.HandleFunc(prefix+"/attributes/{id}", a.PutAttribute).Methods("PUT")

	a.router.HandleFunc(prefix+"/attributes_values", a.GetAttributes_values).Methods("GET")
	a.router.HandleFunc(prefix+"/attributes_values/{id}", a.DeleteAttribute_valueById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/attributes_values", a.PostAttribute_value).Methods("POST")
	a.router.HandleFunc(prefix+"/attributes_values/{id}", a.PutAttribute_value).Methods("PUT")

	a.router.HandleFunc(prefix+"/attributes_values_products", a.GetAettributes_values_products).Methods("GET")
	a.router.HandleFunc(prefix+"/attributes_values_products", a.DeleteAttributes_values_productsById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/attributes_values_products", a.PostAttributes_values_products).Methods("POST")

	a.router.HandleFunc(prefix+"/categories_products", a.GetCategories_products).Methods("GET")
	a.router.HandleFunc(prefix+"/categories_products", a.DeleteCategories_productsById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/categories_products", a.PostCategories_products).Methods("POST")
}

func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
