package controllers

import (
	"fmt"
	"net/http"
)

func RenderEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Render an employee")
	return
}