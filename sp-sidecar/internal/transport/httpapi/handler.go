package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"sp-sidecar/internal/application"
	"sp-sidecar/internal/domain"
)


type DocumentHandler struct {
    service *application.DocumentSourceService
}

func NewDocumentHandler(s *application.DocumentSourceService)(*DocumentHandler){
    return &DocumentHandler{
        service: s,
    }
}

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


func (h *DocumentHandler)ListDocuments(w http.ResponseWriter,r *http.Request){
    
    type documentResponse struct{
        Status string `json:"status"`
        Documents []domain.Document `json:"documents"`
    }

    w.Header().Set("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    docs := h.service.ListDocuments()
    
    response := documentResponse{
        Status: "ok",
        Documents: docs,
    }

    errJsonEncoder := json.NewEncoder(w).Encode(response)
    if errJsonEncoder != nil {
        log.Printf("invalid encoder: %v", errJsonEncoder.Error())
    }
    
}