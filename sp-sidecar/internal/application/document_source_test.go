package application

import (
	
	"slices"
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
        {
            Name: "Mihawk",
            TenantID: "Bank-Cross-Guild",
            Groups: []string{
                "executive",
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
            idWanted []string
            want int
    }{
        {
            user: users[0],
            idWanted: []string{"doc-2"},
            want: 1,
        },
        {
            user: users[1],
            idWanted: []string{
                "doc-1",
                "doc-2",
            },
            want: 2,
        },
        {
            user: users[2],
            idWanted: []string{
                "doc-1",
                "doc-2",
                "doc-4",
            },
            want: 3,
        },
        {
            user: users[3],
            idWanted: []string{},
            want: 0,
        },
        {
            user :users[4],
            idWanted: []string{
                "doc-3",
            },
            want: 1,
        },
    }

    repository := sharepoint.NewMockSharePointSource(docs)
    service := NewDocumentSourceService(repository)
    
    for _,test := range tests{
        
        
        t.Run(test.user.Name,func(t *testing.T) {
            var got []domain.Document
            var IDs []string
            got = service.ListAccessibleDocuments(test.user)
            for _, g :=range got{
                if ! g.CanAccess(test.user) {
                    t.Fatalf("returned inaccessible document: %q", g.ID)
                }
            }
            if len(got) != test.want {
                t.Errorf("excepted: %d, got %d",test.want, len(got))
            }
            for _,g := range got{
                IDs = append(IDs, g.ID)
            }
    
            
            if slices.Compare(IDs,test.idWanted) != 0 {
                t.Fatalf("expected: %v, got: %v",test.idWanted, IDs)
            }
        })
        

    }
}