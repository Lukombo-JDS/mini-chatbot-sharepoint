package httpapi

import "net/http"


func Router()(*http.ServeMux){
    
    serveMux := http.NewServeMux()
    serveMux.HandleFunc("GET /healthz", HealthHandler)
    
    return serveMux
}