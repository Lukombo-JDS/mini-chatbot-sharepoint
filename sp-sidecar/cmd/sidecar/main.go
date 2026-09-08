package main

import (
	"fmt"
	"log"
	"net/http"
	"sp-sidecar/internal/transport/httpapi"
)

func main(){

    server := &http.Server{
        Addr: ":8080",
        Handler: httpapi.Router(),
    }

    fmt.Println("Server started at 8080...")
    err:=server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}