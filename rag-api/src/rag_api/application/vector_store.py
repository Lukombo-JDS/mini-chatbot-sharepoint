from typing import Protocol
from src.rag_api.domain.models import Chunk


class VectorStore(Protocol):

    def add(self, chunks: list[Chunk], embeddings: list[list[float]]) -> None:
        ...

    def search(self, query_embedding: list[float], top_k: int) -> list[Chunk]:
        ...