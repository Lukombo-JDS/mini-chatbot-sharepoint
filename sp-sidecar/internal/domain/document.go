package domain

import (
	"errors"
	"slices"
)

var ErrCollectionEmpty = errors.New("collection empty")
var ErrDocumentNotFound = errors.New("document not found")

type Document struct{
    ID string
    Title string
    Content string
    TenantID string
    AllowedGroups  []string
}

type User struct{
    Name string
    TenantID string
    Groups []string
}


func (d Document)CanAccess(user User)(bool){

    if user.TenantID != d.TenantID {
        return false
    }

    for _,userGroup := range user.Groups{
        if slices.Contains(d.AllowedGroups,userGroup){
            return true
        }
    }

    return false
}