package api

import (
	"net/http"

	"GitHab/Standart_Server_API/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

type StoreToken struct {
	Token string
}

func (api *API) Start() error {
	if err := api.configreLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	api.configreRouterField()
	if err := api.configreStorageField(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.BindAddr, api.router)
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
