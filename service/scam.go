package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zinkworks/grad/server/repository"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Println(json.NewEncoder(w).Encode("service is online"))
}

func ListReport(w http.ResponseWriter, r *http.Request) {
	results := repository.GetALLReport()
	fmt.Println(json.NewEncoder(w).Encode(results))
}
