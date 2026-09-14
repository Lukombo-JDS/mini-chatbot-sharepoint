from typing import Protocol

from src.rag_api.domain.models import Document, User


class DocumentSource(Protocol):
    
    def list_documents(self, user: User) -> list[Document]:
        ...