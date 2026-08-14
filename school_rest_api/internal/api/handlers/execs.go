package handlers

import (
	"fmt"
	"net/http"
)

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	msg := fmt.Sprintf("Hello %s Execs Route", r.Method)
	w.Write([]byte(msg))
	fmt.Println(msg)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(msg))
		fmt.Println(msg)
	case http.MethodPost:
		fmt.Println("Query Parameters:", r.URL.Query())
		fmt.Println("name:", r.URL.Query().Get("name"))

		//parse form data
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing form data:", err)
			return
		}

		fmt.Println("Form from POST methods: ", r.Form)

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
