from src.rag_api.domain.models import Document, User
from dataclasses import dataclass
import httpx

@dataclass
class SharepointClient:
    documents_url: str

    def list_documents(self, user: User) -> list[Document]:

        groups = ",".join(user.groups)

        headers = {
            "X-User-ID": user.user_id,
            "X-Tenant-ID": user.tenant_id,
            "X-Groups": groups
        }

        response = httpx.get(
            url=self.documents_url,
            headers=headers,
            timeout=10.0
        )

        if response.raise_for_status() is not None:
            
            results = response.json()
            docs = results["documents"]
    
            docs_received: list[Document] = []
    
            for d in docs:
                docs_received.append(
                    Document(
                        id=d["ID"],
                        title=d["Title"],
                        content=d["Content"],
                        tenant_id=d["TenantID"],
                        allowed_groups=d["AllowedGroups"]
                    )
                )
            
            return docs_received
        else:
            return []
