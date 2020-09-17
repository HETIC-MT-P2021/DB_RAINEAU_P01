package controllers

import (
	"github.com/gorilla/mux"
	"gobdd/database"
	"gobdd/helpers"
	"gobdd/models"
	"log"
	"net/http"
)

func RenderOrder(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	muxVars := mux.Vars(r)

	strID := muxVars["id"]

	orderID, err := helpers.ParseUInt64(strID)

	if err != nil {
		log.Printf("could not parse int: %v", err)
		return
	}

	orderProducts, err := repository.GetOrderProductsById(orderID)
	if err != nil {
		log.Printf("could not get customer data: %v", err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, orderProducts)

	return
}