package domain

import "errors"

var ErrCollectionEmpty = errors.New("collection empty")

type Document struct{
    ID string
    Title string
    Content string
    TenantID string
    AllowedGroups  []string
}