package application

import (
	"sp-sidecar/internal/domain"
)

type DocumentSourceService struct{
    repository DocumentSource
}


func NewDocumentSourceService(repository DocumentSource)(*DocumentSourceService){
    return &DocumentSourceService{
        repository: repository,
    }
}

func (ds *DocumentSourceService)ListDocuments()([]domain.Document){
    return ds.repository.ListDocuments()
}

func (ds *DocumentSourceService)GetDocument(id string)(domain.Document,error){
    return ds.repository.GetDocument(id)
}