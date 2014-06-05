package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RestAPI struct {
	router   *http.ServeMux
	provider Provider
}

func NewRestAPI(p Provider) *RestAPI {
	api := &RestAPI{
		router:   http.NewServeMux(),
		provider: p,
	}
	api.registerEndPoints()
	return api
}

func (api *RestAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin,Authorization,Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		api.router.ServeHTTP(w, r)
	}
}

func (api *RestAPI) registerEndPoints() {
	api.router.HandleFunc("/ping", api.handlePing)
	api.router.HandleFunc("/overlay", api.handleGetOverlays)
}

func (api *RestAPI) handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *RestAPI) handleGetOverlays(w http.ResponseWriter, r *http.Request) {
	overlays := api.provider.Overlays()
	b, err := json.Marshal(overlays)
	if err != nil {
		log.Printf("Error encoding overlays: %q", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(b)
	}
}
