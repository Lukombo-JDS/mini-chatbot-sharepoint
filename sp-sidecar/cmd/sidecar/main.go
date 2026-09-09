package main

import (
	"fmt"
	"log"
	"net/http"
	"sp-sidecar/internal/application"
	"sp-sidecar/internal/infrastructure/sharepoint"
	"sp-sidecar/internal/transport/httpapi"
)

func main(){

    repository := sharepoint.NewMockSharePointSource(sharepoint.CollectionMockSharePointSource())

    service := application.NewDocumentSourceService(repository)
    

    router := httpapi.Router(
        httpapi.NewDocumentHandler(service),
    )

    server := &http.Server{
        Addr: ":8080",
        Handler: router,
    }

    fmt.Println("Server started at 8080...")
    err:=server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}