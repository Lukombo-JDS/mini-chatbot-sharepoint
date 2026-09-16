from dataclasses import dataclass
from src.rag_api.application.document_source import DocumentSource
from src.rag_api.domain.models import Chunk, Document, User


@dataclass
class DocumentRetrievalService:
    document_source: DocumentSource

    def get_documents(self, user:User) -> list[Document]:
        return self.document_source.list_documents(user)

@dataclass
class ChunkingService:

    def chunk_document(self, document: Document, chunk_size:int) -> list[Chunk]:

        if chunk_size <=0 :
            return []

        
        count_chunk: int = 0
        chunks: list[Chunk] = []
        words: list[str] = document.content.split()
        chunk_content: list[str] = []

        for wrd in words:

            if len(chunk_content) < chunk_size:
                chunk_content.append(wrd)
                
            else:
                chunks.append(
                    Chunk(
                        id=document.id+":chunk-"+str(count_chunk),
                        document_id=document.id,
                        chunk_index=count_chunk,
                        tenant_id=document.tenant_id,
                        content= " ".join(chunk_content),
                        allowed_groups=document.allowed_groups,
                    )
                )
                
                count_chunk += 1
                chunk_content = []
                chunk_content.append(wrd)
                
        if len(chunk_content) > 0:
            chunks.append(
                Chunk(
                    id=document.id+":chunk-"+str(count_chunk),
                    document_id=document.id,
                    chunk_index=count_chunk,
                    tenant_id=document.tenant_id,
                    content= " ".join(chunk_content),
                    allowed_groups=document.allowed_groups,
                )
            )

                
        return chunks