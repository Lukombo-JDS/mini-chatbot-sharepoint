package sharepoint

import (
	"fmt"
	"slices"
	"sp-sidecar/internal/domain"
)

//Mock of a SharePoint source of documents
type MockSharePointSource struct{
    documents []domain.Document
}

//Function creating a new SharePoint Source
func NewMockSharePointSource(collection []domain.Document)(*MockSharePointSource){
        return &MockSharePointSource{
            documents: collection,
        }
}

//Collection of mock SharePoint Documents source
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


/*List Document a user can access */
func (ms *MockSharePointSource)ListAccessibleDocuments(user domain.User)([]domain.Document){

    var documents []domain.Document

    for _,d :=range ms.documents {
       
       if d.CanAccess(user) {
           documents = append(documents, d)
       }
    }
    
    return slices.Clone(documents)
}

//Implementation of method listing documents
func (dsr *MockSharePointSource)ListDocuments()([]domain.Document){

    var documents []domain.Document
    
    for _,doc :=range dsr.documents {
        documents = append(documents, doc)
    }

    if len(documents) == 0 {
        return nil
    }

    return slices.Clone(documents)
    
}

//Implementation of method getting document by ID
func (dsr *MockSharePointSource)GetDocument(id string)(domain.Document,error){


        for _,doc := range dsr.documents {
            if doc.ID == id {
                return doc,nil
            }
        }

        return domain.Document{}, fmt.Errorf(
            "%w: document id = %q",
            domain.ErrDocumentNotFound,
            id,
        )
}




