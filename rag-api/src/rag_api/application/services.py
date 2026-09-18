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

    def chunk_document(self, document: Document, chunk_size:int, overlap:int) -> list[Chunk]:

        if chunk_size <= 0 :
            raise ValueError("invalid chunk size")

        if overlap < 0 :
            raise ValueError("invalid overlap value")

        if overlap >= chunk_size:
            raise ValueError("invalid overlap equal chunk size")
        
        end: int = chunk_size
        chunks: list[Chunk] = []
        words: list[str] = document.content.split()
        chunk_content: list[str] = []
        count_chunk:int = 0

        start = 0
        
        while start < len(words):

            end = min(start+chunk_size,len(words))
        
            chunk_content = words[start:end]

            chunks.append(
                Chunk(
                    id = document.id + ":chunk-" + str(count_chunk),
                    document_id = document.id,
                    chunk_index = count_chunk,
                    tenant_id = document.tenant_id,
                    content = " ".join(chunk_content),
                    allowed_groups = document.allowed_groups,
                )
            )

            if end == len(words):
                break

            count_chunk += 1 # increment the id
            start += (chunk_size-overlap) # sliding window

                
        return chunks