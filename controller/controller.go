package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"zinkworks/grad/server/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleReqeust() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	apiRoutes := myRouter.PathPrefix("/api/v1").Subrouter()
	myRouter.HandleFunc("/health", service.Health).Methods("GET")
	apiRoutes.HandleFunc("/Report", service.ListReport).Methods("GET")
	return myRouter
}

func Run(routes *mux.Router) {
	host := fmt.Sprintf("%v", os.Getenv("host"))
	fmt.Println("host", host)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{host},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := c.Handler(routes)
	log.Fatal("server has started", http.ListenAndServe(":8080", handler))
}
