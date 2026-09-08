package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request){

    type HealthResponse struct{
        Status string `json:"status"`
    }
        
    w.Header().Set("content-type","application/json")
    w.WriteHeader(http.StatusOK)
    response := HealthResponse{
        Status: "ok",
    }

    errJsonEncoder := json.NewEncoder(w).Encode(response)
    if errJsonEncoder != nil {
        log.Printf("invalid encoder: %v", errJsonEncoder.Error())
    }

}