package handler

import (
	"fmt"
	"net/http"
)

func UserDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Service 2 UserDelete")

	w.WriteHeader(http.StatusInternalServerError)
}
