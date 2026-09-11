package httpapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sp-sidecar/internal/application"
	"sp-sidecar/internal/domain"
	"strings"
)

var ErrUnauthorized = errors.New("missing user identity")

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


func (h *DocumentHandler)ListDocuments(w http.ResponseWriter, r *http.Request){
    
    type documentResponse struct{
        Status string `json:"status"`
        Documents []domain.Document `json:"documents"`
    }

    type documentErrorResponse struct{
        Error string `json:"error"`
    }

    user,err := userFormRequest(r)

    w.Header().Set("content-type", "application/json")
    

    if errors.Is(err, ErrUnauthorized) {
        
        w.WriteHeader(http.StatusUnauthorized)

    
        response := documentErrorResponse{
            Error: ErrUnauthorized.Error(),
        }

        errJsonEncoder := json.NewEncoder(w).Encode(response)
        if errJsonEncoder != nil {
            log.Printf("invalid encoder: %v", errJsonEncoder.Error())
        }
        return 
        
    }
        
    docs := h.service.ListAccessibleDocuments(user)

    w.WriteHeader(http.StatusOK)
    
    response := documentResponse{
        Status: "ok",
        Documents: docs,
    }

    errJsonEncoder := json.NewEncoder(w).Encode(response)
    if errJsonEncoder != nil {
        log.Printf("invalid encoder: %v", errJsonEncoder.Error())
    }
    
}

func userFormRequest(r *http.Request)(domain.User,error){

    userID:=r.Header.Get("X-User-ID")
    tenantID:=r.Header.Get("X-Tenant-ID")
    groups:=r.Header.Get("X-Groups")

    if userID == "" || tenantID == "" {
        return domain.User{}, ErrUnauthorized
    }

    if len(groups) == 0 {
        return domain.User{
            Name: userID,
            TenantID: tenantID,
            Groups: []string{},
        },nil
    }
    

    g := strings.Split(groups,",")

    var Groups []string
    
    for _,elmt := range g {
        
        word:=strings.TrimSpace(elmt)

        Groups = append(Groups, word)
    }

    return domain.User{
        Name: userID,
        TenantID: tenantID,
        Groups: Groups,
    },nil
    
}