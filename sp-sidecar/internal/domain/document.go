package domain

import "errors"

var ErrCollectionEmpty = errors.New("collection empty")
var ErrDocumentNotFound = errors.New("document not found")

type Document struct{
    ID string
    Title string
    Content string
    TenantID string
    AllowedGroups  []string
}