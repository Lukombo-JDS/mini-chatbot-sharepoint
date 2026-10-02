import pytest

from src.rag_api.domain.models import Chunk
from src.rag_api.infrastructure.vector_store.in_memory import InMemoryVectorStore


chunks = [
    Chunk(
        id="chunk-1",
        document_id="doc-1:chunk-1",
        chunk_index=0,
        content="Luffy the CEO of StrawHats Banks",
        tenant_id="StrawHats Banks",
        allowed_groups=["executive", "employee"]
    ),
    Chunk(
        id="chunk-2",
        document_id="doc-2:chunk-2",
        chunk_index=1,
        content="Nami the new CFO",
        tenant_id="StrawHats Banks",
        allowed_groups=["executive", "employee", "risk"]
    ),
]

embeddings = [
    [0.2, -0.4, 0.5, -0.7],
    [0.7, -0.8, -0.5, 0.9]
]

@pytest.mark.parametrize(
    "indice",
    [
        0,
        1
    ]
)
def test_check_add(indice):

    InMemoryVectorStore().add(chunks,embeddings)

    store = InMemoryVectorStore().memory
    InMemoryVectorStore().memory = []
    assert chunks[indice] == store[indice].chunk
    assert embeddings[indice] == store[indice].embedding


def test_check_second_call():

    chunks_2 = [Chunk(
        id = "chunk-3",
        document_id="doc-3:chunk-3",
        chunk_index=3,
        tenant_id="Kara Bank",
        content="Luffy burrow 500€",
        allowed_groups=["client", "employee", "executive"]
    ),
        Chunk(
            id = "chunk-4",
            document_id="doc-4:chunk-4",
            chunk_index=4,
            tenant_id="Kara Bank",
            content="Zoro became COO",
            allowed_groups=["risk", "executive"]
        )
    ]

    embeddings_2 = [
        [-0.3, 0.6, 9.0, -0.4],
        [0.8, 0.1, -0.1, -0.7]
    ]

    # InMemoryVectorStore().add(chunks, embeddings)

    InMemoryVectorStore().add(chunks_2, embeddings_2)

    store = InMemoryVectorStore().memory
    
    print("store: ", store)
    print("store[3]: ", store[3].chunk)
    print("store[4]: ", store[4].chunk)
        
    assert chunks[0] == store[0].chunk
    assert embeddings[0] == store[0].embedding

    assert chunks[1] == store[1].chunk
    assert embeddings[1] == store[1].embedding


def test_entry_error():

    with pytest.raises(expected_exception=ValueError, match="invalid entries"):

        InMemoryVectorStore().add([chunks[0]], embeddings)