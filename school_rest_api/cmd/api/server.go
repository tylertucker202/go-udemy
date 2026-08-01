package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	//teachers/{id}
	fmt.Println(r.Method)
	msg := fmt.Sprintf("Hello %s Teachers Route", r.Method)
	w.Write([]byte(msg))
	fmt.Println(msg)
	switch r.Method {
	case http.MethodGet:
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")
		queryParams := r.URL.Query()
		fmt.Println("User ID:", userID)
		fmt.Println("Query name:", queryParams["name"])
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")

		if sortorder == "" {
			sortorder = "asc"
		}

		fmt.Println("Sort by:", sortby)
		fmt.Println("Key:", key)
		fmt.Println("Sort order:", sortorder)

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

func studentsHandler(w http.ResponseWriter, r *http.Request) {
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

func execsHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {

	port := ":3000"

	http.HandleFunc("/root", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
