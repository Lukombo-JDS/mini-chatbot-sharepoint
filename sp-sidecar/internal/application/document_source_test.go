package application

import (
	"sp-sidecar/internal/domain"
	"sp-sidecar/internal/infrastructure/sharepoint"
	"testing"
)

func TestListAccessibleDocuments(t *testing.T){
    //Arrange

    users := []domain.User{
        {
            Name: "Luffy",
            TenantID: "Bank-Kara",
            Groups: []string{
                    "clients",
            },
        },
        {
            Name: "Zoro",
            TenantID: "Bank-Kara",
            Groups: []string{
                "employee",
            },
        },
        {
            Name: "Nami",
            TenantID: "Bank-Kara",
            Groups: []string{
                    "risk",
                    "executive",
                },
        },
        {
            Name: "Baggy",
            TenantID: "Bank-Cross-Guild",
            Groups: []string{
                    "trainee",
                },
        },
    }

    docs := []domain.Document{
        {
            ID: "doc-1",
            Title: "bank-Kara HQ",
            Content: "The organisation of the Bank of Kara...",
            TenantID: "Bank-Kara",
            AllowedGroups: []string{
                "employee",
                "executive",
            },
        },
        {
            ID: "doc-2",
            Title: "Bank-Kara ROI",
            Content: "The ROI of this year is 200 billion § ...",
            TenantID: "Bank-Kara",
            AllowedGroups: []string{
                "executive",
                "risk",
                "employee",
                "clients",
            },
        },
        {
            ID: "doc-3",
            Title: "StrawHats Money",
            Content: "Revenue: 300 000§",
            TenantID: "Bank-Cross-Guild",
            AllowedGroups: []string{
                "clients",
                "executive",
            },
        },
        {
            ID: "doc-4",
            Title: "Bank Kara 4 years Plans",
            Content: "The plans for the 4 futur years with Cross Guild...",
            TenantID: "Bank-Kara",
            AllowedGroups: []string{
                "risk",
                "executive",
            },
        },
    }
    
    tests := []struct {
            user domain.User
            // document domain.Document
            want int
    }{
        {
            user: users[0],
            // document: docs[0],
            want: 1,
        },
        {
            user: users[1],
            // document: docs[1],
            want: 2,
        },
        {
            user: users[2],
            // document: docs[2],
            want: 3,
        },
        {
            user: users[3],
            // document: docs[3],
            want: 0,
        },
    }

    repository := sharepoint.NewMockSharePointSource(docs)
    service := NewDocumentSourceService(repository)
    for _,test := range tests{
        var got []domain.Document
        t.Run(test.user.Name,func(t *testing.T) {
            got = service.ListAccessibleDocuments(test.user)
        })
        if len(got) != test.want {
            t.Errorf("excepted: %d, got %d",test.want, len(got))
        }
    }
}