from dataclasses import dataclass
from src.rag_api.domain.models import Chunk
import numpy as np

@dataclass
class VectorEntry:
    chunk: Chunk
    embedding: list[float]


class InMemoryVectorStore:
    
    memory: list[VectorEntry] = []

    def add(self, chunks: list[Chunk], embeddings: list[list[float]]):

        if len(chunks) != len(embeddings):
            raise ValueError("invalid entries")

        for chunk,embedding in zip(chunks,embeddings, strict=True):

            self.memory.append(VectorEntry(
                chunk=chunk,
                embedding=embedding
            ))
      
    def search(self, query_embedding: list[float], top_k: int) -> list[Chunk]:

       query_embedding_array = np.array(query_embedding)

       store = self.memory
       retrieved_chunks:list[Chunk] = []

       for s in store:

           embedding_array = np.array(s.embedding)

           cosine_similarity = np.dot(
               query_embedding_array, 
               embedding_array) / np.dot(
               np.abs(query_embedding_array),
               np.abs(embedding_array)
           )


           if cosine_similarity >= 0.8 and len(retrieved_chunks) <= top_k:

               retrieved_chunks.append(s.chunk)

       return retrieved_chunks