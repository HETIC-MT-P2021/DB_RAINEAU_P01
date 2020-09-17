package controllers

import (
	"fmt"
	"net/http"
)

func RenderShop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Render a shop")
	return
}