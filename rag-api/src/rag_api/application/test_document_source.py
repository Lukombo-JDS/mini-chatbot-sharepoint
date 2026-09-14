from src.rag_api.application.services import DocumentRetrievalService
from src.rag_api.domain.models import Document, User

class FakeDocumentSource:

    def list_documents(self, user: User) -> list[Document]:
        return [Document(
            id="doc-1",
            title="organigramme",
            tenant_id="Bank-Cross-Guild",
            content="Crocodile, Mihawk les bossent",
            allowed_groups=["executive"]
        ),
        Document(
            id="doc-2",
            title="Fortunes",
            tenant_id="Bank-Cross-Guild",
            content="Baggy le riche...",
            allowed_groups=[
                "risk",
                "executive"
            ]
        )
        ]

def test_document_retrieval_service():

    user = User(
        user_id="Nami",
        tenant_id="Bank-Cross-Guild",
        groups=["risk"]
    )

    documents_retrieved = DocumentRetrievalService(FakeDocumentSource()).get_documents(user=user)

    assert len(documents_retrieved) == 2
    assert documents_retrieved[0].id == "doc-1"
    assert documents_retrieved[1].id == "doc-2"