package application

import "sp-sidecar/internal/domain"

type DocumentSource interface{
    ListDocuments()([]domain.Document, error)
    GetDocument(id string)(domain.Document,error)
}