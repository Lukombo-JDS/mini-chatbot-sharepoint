package main

import (
	"fmt"
	"sp-sidecar/internal/domain"
)

func main(){

    doc_employee:=domain.Document{
        ID: "doc-1",
        Name: "doc-employee",
        Content: "Luffy is the captain",
        TenantID: "bank-a",
        AllowGroups: []string{
            "employee",
        },
    }

    doc_risk:=domain.Document{
        ID: "doc-2",
        Name: "doc-risk",
        Content: "ROI this year: 20 000 000€",
        TenantID: "bank-a",
        AllowGroups: []string{
            "risk",
        },
    }

    doc_executive:=domain.Document{
        ID: "doc-3",
        Name: "doc-executive",
        Content: "The CFO is Nami",
        TenantID: "bank-a",
        AllowGroups:[]string{
            "executive",
        },
    }

    fmt.Printf("doc_employee: %v\n", doc_employee)
    fmt.Printf("doc_risk: %v\n", doc_risk)
    fmt.Printf("doc_executive: %v\n", doc_executive)

    
}