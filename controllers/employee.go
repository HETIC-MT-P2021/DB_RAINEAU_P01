package controllers

import (
	"gobdd/database"
	"gobdd/helpers"
	"gobdd/models"
	"log"
	"net/http"
)

func RenderEmployees(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	employeesData, err := repository.GetEmployees()
	if err != nil {
		log.Printf("could not get employees data: %v", err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, employeesData)

	return
}