package httpapi

import "net/http"


func Router(dh *DocumentHandler)(*http.ServeMux){

    
    
    serveMux := http.NewServeMux()
    serveMux.HandleFunc("GET /healthz", HealthHandler)
    serveMux.HandleFunc("GET /documents", dh.ListDocuments)
    
    return serveMux
}