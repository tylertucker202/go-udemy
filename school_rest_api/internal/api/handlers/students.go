package handlers

import (
	"fmt"
	"net/http"
)

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	msg := fmt.Sprintf("Hello %s Students Route", r.Method)
	w.Write([]byte(msg))
	fmt.Println(msg)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(msg))
		fmt.Println(msg)
	case http.MethodPost:
		w.Write([]byte(msg))
		fmt.Println(msg)
	case http.MethodPut:
		w.Write([]byte(msg))
		fmt.Println(msg)
	case http.MethodPatch:
		w.Write([]byte(msg))
		fmt.Println(msg)
	case http.MethodDelete:
		w.Write([]byte(msg))
		fmt.Println(msg)
	}
}
