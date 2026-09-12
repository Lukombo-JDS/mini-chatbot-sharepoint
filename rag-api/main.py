from src.rag_api.domain.models import User
from src.rag_api.infrastructure.sharepoint_client import SharepointClient


def main():

    docs = SharepointClient("http://localhost:8080/documents").list_documents(
        User(
            user_id="Luffy",
            tenant_id="Bank-Kara",
            groups=["risk"]
        )
    )

    print(docs)


if __name__ == "__main__":
    main()
