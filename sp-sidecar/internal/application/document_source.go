package application

import (
	"slices"
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

func (ds *DocumentSourceService)ListAccessibleDocuments(user domain.User)([]domain.Document){

        docsAllowed := []domain.Document{}
        
        for _,d := range ds.repository.ListDocuments() {

            if d.CanAccess(user) {
                docsAllowed = append(docsAllowed, d)
            }
        }

        return slices.Clone(docsAllowed)
}