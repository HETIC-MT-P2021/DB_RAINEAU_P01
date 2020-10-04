package controllers

import (
	"github.com/gorilla/mux"
	"gobdd/database"
	"gobdd/helpers"
	"gobdd/models"
	"log"
	"net/http"
)

func RenderOffice(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	muxVars := mux.Vars(r)

	strID := muxVars["id"]

	officeID, err := helpers.ParseUInt64(strID)

	if err != nil {
		log.Printf("could not parse int: %v", err)
		return
	}

	officeData, err := repository.GetOfficeByID(officeID)
	if err != nil {
		log.Printf("could not get office data: %v", err)
		return
	}

	employees, err := repository.GetOfficeEmployees(officeData)
	if err != nil {
		log.Printf("could not get orders data: %v", err)
		return
	}

	officeData.Employees = employees

	helpers.WriteJSON(w, http.StatusOK, officeData)

	return
}