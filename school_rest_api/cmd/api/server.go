package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := sqlconnect.ConnectDb()

	db.Close()

	if err != nil {
		fmt.Println("Error=====: ", err)
	}

	port := os.Getenv("API_PORT")
	cert := "cert.pem"
	key := "key.pem"

	// mux := http.NewServeMux()

	tlsCofig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := middlewares.NewRateLimiter(5, time.Minute)

	// hppOptions := middlewares.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := middlewares.Cors(rl.Middleware(middlewares.ResponseTimeMiddleware(middlewares.SecurityHeaders(middlewares.CompressionMiddleware(middlewares.Hpp(hppOptions)(mux))))))
	// secureMux := utils.ApplyMiddlewares(mux,
	// 	middlewares.CompressionMiddleware,
	// 	middlewares.SecurityHeaders,
	// 	middlewares.ResponseTimeMiddleware,
	// 	rl.Middleware,
	// 	middlewares.Hpp(hppOptions),
	// 	middlewares.Cors)

	router := router.Router()
	secureMux := utils.ApplyMiddlewares(router, middlewares.SecurityHeaders)

	server := &http.Server{
		Addr: port,
		//Handler:   middlewares.SecurityHeaders(mux),
		//Handler:   middlewares.Cors(mux),
		Handler:   secureMux,
		TLSConfig: tlsCofig,
	}

	fmt.Println("Server is running on port: ", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
