package application

import "sp-sidecar/internal/domain"

type DocumentSource interface{
    ListDocuments()([]domain.Document)
    GetDocument(id string)(domain.Document,error)
}