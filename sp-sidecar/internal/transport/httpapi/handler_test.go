package httpapi

import (
	"encoding/json/v2"
	"io"
	"net/http"
	"net/http/httptest"

	"slices"
	"sp-sidecar/internal/application"
	"sp-sidecar/internal/domain"
	"sp-sidecar/internal/infrastructure/sharepoint"
	"strings"
	"testing"
)


func injestionDependanciesDocument()(*application.DocumentSourceService){

    collection := sharepoint.CollectionMockSharePointSource()

    repository := sharepoint.NewMockSharePointSource(collection)

    service := application.NewDocumentSourceService(repository)

    return service
    
}

//Testing the call of the Document accessible from
// the request of a presume valid user into the header
// I assert the status code for the user and the matching 
// documents IDs

func TestListDocumentsHandler(t *testing.T) {

    // Arrange
    tests := []struct{
            Name string
            Headers domain.User
            ExpectedStatus int
            ExpectedDocumentsIDs []string
    }{
        {
            Name: "missing_identity",
            Headers: domain.User{},
            ExpectedStatus: http.StatusUnauthorized,
            ExpectedDocumentsIDs: []string{},
        },
        {
            Name: "nami_bank_kara",
            Headers: domain.User{
                Name: "Nami",
                TenantID: "Bank-Kara",
                Groups: []string{
                    "risk",
                    "executive",
                },
            },
            ExpectedStatus: http.StatusOK,
            ExpectedDocumentsIDs: []string{
                "doc-1",
                "doc-3",
                "doc-4",
            },
        },
        {
            Name: "mihawk_cross_tenant",
            Headers: domain.User{
                Name: "Mihawk",
                TenantID: "Bank-Cross-Guild",
                Groups: []string{
                    "risk",
                    "executive",
                },
            },
            ExpectedStatus: http.StatusOK,
            ExpectedDocumentsIDs: []string{
                "doc-2",
            },
        },
        {
            Name: "missing_groups",
            Headers: domain.User{
                Name: "Luffy",
                TenantID: "Bank-Kara",
                Groups: []string{},
            },
            ExpectedStatus: http.StatusOK,
            ExpectedDocumentsIDs: []string{},
        },
    }

    type requestHeader struct{
        UserID string 
        TenantID string 
        Groups []string
    }

    type documentHandlerResponse struct{
        Status string `json:"status"`
        Documents []domain.Document `json:"documents"`
    }

    type documentHandlerError struct{
        Error string `json:"error"`
    }
    service := injestionDependanciesDocument()

    handler := NewDocumentHandler(service)

    // Act
    for _,test := range tests{

        t.Run(test.Name, func(t *testing.T) {
            
            
            request := httptest.NewRequest("GET", "/documents", /*bytes.NewReader(res)*/nil)

            //sending user identity
            request.Header.Set("X-User-ID",test.Headers.Name)
            request.Header.Set("X-Tenant-ID", test.Headers.TenantID)
            request.Header.Set("X-Groups", strings.Join(
                test.Headers.Groups,
                ",",
            ))
            
            recorder := httptest.NewRecorder() //Create a new recorder
        
            handler.ListDocuments(recorder,request) // Calling the handler

            statusCode := recorder.Code
            
            res := recorder.Result() // saving the result
            defer res.Body.Close()

            data,err := io.ReadAll(res.Body)
            if err != nil {
                t.Fatalf("invalid read body: %v\n", err)
            }

            var documentHandlerResponse documentHandlerResponse
            var handlerError documentHandlerError
            var DocumentIDsClone []string
            var DocumentIDsExpectedClone []string
            var DocumentIDs []string

            t.Logf("server response: %v\n\n", string(data))

            errUnmarshall := json.Unmarshal(data, &documentHandlerResponse)
            if errUnmarshall != nil {
                t.Fatalf("invalid unmarshall: %v\n", errUnmarshall.Error())
            }

            //construction of the string of documents IDs receive from
            // the memory based and the user request
            for _,d :=range documentHandlerResponse.Documents {
                DocumentIDs = append(DocumentIDs, d.ID)
            }

            //cloning before sorting to not change the concrete slices

            DocumentIDsClone = slices.Clone(DocumentIDs)
            DocumentIDsExpectedClone = slices.Clone(test.ExpectedDocumentsIDs)

            //sorting both slices to compare in the same order element by element.
            slices.Sort(DocumentIDsClone)
            slices.Sort(test.ExpectedDocumentsIDs)
        
            //Assert

            //check the status code: CanAccess is working or not
            // to let the right user access the data
            if statusCode != test.ExpectedStatus {
                t.Errorf("expected %d, got %d",test.ExpectedStatus, statusCode)

                if statusCode == http.StatusUnauthorized {
                    err = json.Unmarshal(data, &handlerError)
                    if err != nil {
                        t.Fatalf("invalid unmarshall: %v\n", err)
                    }

                    if handlerError.Error != ErrUnauthorized.Error() {
                        t.Errorf("expected %v, got %v \n", ErrUnauthorized.Error(), handlerError.Error)
                    }
                }
            }

            //check the receives docs are the accessbile ones based on 
            // the allowed groups and the TenantID
            if slices.Compare(
                DocumentIDsClone,
                DocumentIDsExpectedClone,
            ) != 0 {
                t.Errorf(
                    "expected %v, got %v",
                    DocumentIDsExpectedClone,
                    DocumentIDsClone,
                )
            }
        })
    }
}