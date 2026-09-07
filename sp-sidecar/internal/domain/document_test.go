package domain

import (
	// "sp-sidecar/internal/infrastructure/sharepoint"
	"testing"
)

func TestDocumentCanAccess(t *testing.T){
    //Arrange
    // collection := sharepoint.CollectionMockSharePointSource()

    doc := Document{
        ID: "doc-test",
        Title: "Merry Price",
        Content: "The Merry Cost 2 000 000 Berrys",
        TenantID: "bank-a",
        AllowedGroups: []string{
            "executive",
            "risk",
        },
    }

    tests := []struct {
            name string
            tenantID string
            groups []string
            want bool
    }{
        {
            name: "Luffy",
            tenantID: "bank-a",
            groups: []string{
                    "risk",
            },
            want: true,
        },
        {
            name: "Zoro",
            tenantID: "bank-a",
            groups: []string{
                "employee",
            },
            want: false,
        },
        {
            name: "Nami",
            tenantID: "bank-b",
            groups: []string{
                    "risk",
                    "executive",
                },
            want: false,
        },
        {
            name: "Usopp",
            tenantID: "bank-a",
            groups: []string{
                    "employee",
                    "risk",
                },
            want: true,
        },
    }

   for _, test :=range tests{
       var got bool
       t.Run(test.name, func(t *testing.T) {
           got=doc.CanAccessSimpler(test.tenantID,test.groups)
       })
       if got != test.want {
           t.Errorf("excepted %v, got %v",test.want,got)
       }
   }
}

func TestRefactorDocumentCanAccess(t *testing.T){
    //Arrange
    // collection := sharepoint.CollectionMockSharePointSource()

    doc := Document{
        ID: "doc-test",
        Title: "Merry Price",
        Content: "The Merry Cost 2 000 000 Berrys",
        TenantID: "bank-a",
        AllowedGroups: []string{
            "executive",
            "risk",
        },
    }

    tests := []struct {
            user User
            want bool
    }{
        {
            user: User{
                Name: "Luffy",
                TenantID: "bank-a",
                Groups: []string{
                    "risk",
                },
            },
            want: true,
        },
        {
            user: User{

                Name: "Zoro",
                TenantID: "bank-a",
                Groups: []string{
                    "employee",
                },
            },
            want: false,
        },
        {
            user: User{
                Name: "Nami",
                TenantID: "bank-b",
                Groups: []string{
                    "risk",
                    "executive",
                },
            },
            want: false,
        },
        {
            user: User{
                Name: "Usopp",
                TenantID: "bank-a",
                Groups: []string{
                    "employee",
                    "risk",
                },
            },
            want: true,
        },
    }

   for _, test :=range tests{
       var got bool
       t.Run(test.user.Name, func(t *testing.T) {
           got=doc.CanAccess(test.user)
       })
       if got != test.want {
           t.Errorf("excepted %v, got %v",test.want,got)
       }
   }
}
