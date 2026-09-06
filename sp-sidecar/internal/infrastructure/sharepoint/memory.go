package sharepoint

import (
	"fmt"
	"slices"
	"sp-sidecar/internal/application"
	"sp-sidecar/internal/domain"
)


type MockSharePointSource struct{
    Collection []domain.Document
}

type DocumentSourceReporistory struct {
    repository application.DocumentSource
}


func (dsr *MockSharePointSource)ListDocuments()([]domain.Document,error){

    var documents []domain.Document
    // var documents_copy []domain.Document
    
    for _,doc :=range dsr.Collection {
        documents = append(documents, doc)
    }

    if len(documents) == 0 {
        return nil, fmt.Errorf(
            "%w",
            domain.ErrCollectionEmpty,
        )
    }

    return slices.Clone(documents),nil
    
}



func NewMockSharePointSource(collection []domain.Document)(*MockSharePointSource){

        return &MockSharePointSource{
            Collection: collection,
        }
}


func CollectionMockSharePointSource()([]domain.Document){
    return []domain.Document{
                {
                    ID: "doc-1",
                    Title: "doc-employee",
                    Content: "Luffy is the captain",
                    TenantID: "bank-a",
                    AllowedGroups: []string{
                        "employee",
                    },
                },
                {
                    ID: "doc-2",
                    Title: "doc-risk",
                    Content: "ROI this year: 20 000 000€",
                    TenantID: "bank-a",
                    AllowedGroups: []string{
                        "risk",
                    },   
                },
                {
                    ID: "doc-3",
                    Title: "doc-executive",
                    Content: "The CFO is Nami",
                    TenantID: "bank-a",
                    AllowedGroups:[]string{
                        "executive",
                    },
                },
            }
}

