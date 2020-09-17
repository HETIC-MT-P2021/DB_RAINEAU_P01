package controllers

import (
	"github.com/gorilla/mux"
	"gobdd/database"
	"gobdd/helpers"
	"gobdd/models"
	"log"
	"net/http"
)

func RenderCustomer(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	muxVars := mux.Vars(r)

	strID := muxVars["id"]

	customerID, err := helpers.ParseUInt64(strID)

	if err != nil {
		log.Printf("could not parse int: %v", err)
		return
	}

	customerData, err := repository.GetCustomerByID(customerID)
	if err != nil {
		log.Printf("could not get customer data: %v", err)
		return
	}

	orders, err := repository.GetCustomerOrders(customerData)
	if err != nil {
		log.Printf("could not get orders data: %v", err)
		return
	}

	customerData.Orders = orders

	helpers.WriteJSON(w, http.StatusOK, customerData)

	return
}