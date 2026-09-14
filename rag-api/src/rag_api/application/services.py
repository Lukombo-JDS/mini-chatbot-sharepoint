from dataclasses import dataclass
from src.rag_api.application.document_source import DocumentSource
from src.rag_api.domain.models import Document, User


@dataclass
class DocumentRetrievalService():
    document_source: DocumentSource

    def get_documents(self, user:User) -> list[Document]:
        return self.document_source.list_documents(user)