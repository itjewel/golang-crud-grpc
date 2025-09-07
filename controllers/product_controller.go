package controllers

import (
	"encoding/json"
	"golang-crud/models"
	"golang-crud/service"
	"net/http"
)

type ControllerService struct {
	Service *service.ProductService
}

func (pc *ControllerService) GetProductInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := pc.Service.GetProductService(ctx)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func (pc *ControllerService) AddProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	responseValue, err := pc.Service.AddProduct(ctx, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(responseValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
