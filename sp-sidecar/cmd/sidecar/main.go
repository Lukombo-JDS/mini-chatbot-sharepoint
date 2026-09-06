package main

import (
	"errors"
	"fmt"
	"sp-sidecar/internal/application"
	"sp-sidecar/internal/domain"
	"sp-sidecar/internal/infrastructure/sharepoint"
)

func main(){
    
    collection := sharepoint.CollectionMockSharePointSource()

    repository := sharepoint.NewMockSharePointSource(collection)

    service := application.NewDocumentSourceService(repository)

    docs := service.ListDocuments()

    fmt.Println(docs)

    doc,err_get_document := service.GetDocument("doc-2")

    fmt.Println(doc)

    switch{
        case errors.Is(err_get_document,domain.ErrDocumentNotFound):
            fmt.Println(err_get_document)
    }
    
}