from src.rag_api.application.services import ChunkingService
from src.rag_api.domain.models import Document
import pytest


def test_chunk_document():

    # Arrange
    fake_document = Document(
        title="One Piece Live Action",
        id="doc-1",
        content="The One Piece Live Action was made with heart and effort.  ",
        tenant_id="Bank-Kara",
        allowed_groups=["risk", "executive"]
        
    )
    
    chunk_size = 3

    # Act
    chunks = ChunkingService().chunk_document(document=fake_document,chunk_size=chunk_size)
    # print("CHUNKS: ", chunks)

    content:list[str] = []

    # Assert
    for i,chk in enumerate(chunks):
        
        assert len(chk.content.split()) <= 3
        assert chk.document_id == fake_document.id
        assert chk.tenant_id == fake_document.tenant_id
        assert chk.allowed_groups == fake_document.allowed_groups

        content.append(chk.content)

        assert chk.chunk_index == i
        assert chk.id == f"{fake_document.id}:chunk-{i}"
        

    reconstructed_content = " ".join(content)

    assert reconstructed_content.split() == fake_document.content.split()


@pytest.mark.parametrize(
    "content, chunk_size, expected_chunks",
    [
        ("The One Piece", 5, 1),    # content < chunk_size
        ("The One Piece", 3, 1),    # content == chunk_size
        ("The One Piece", 2, 2),    # content > chunk_size
        ("The One Piece", 1, 3)     # 1 word == chunk_size
    ]
)
def test_chunk_sizes(content, chunk_size, expected_chunks):

     # Arrange
     fake_document = Document(
         title="One Piece Live Action",
         id="doc-1",
         content=content,
         tenant_id="Bank-Kara",
         allowed_groups=["risk", "executive"]
         
     )
      
     # Act
     chunks = ChunkingService().chunk_document(document=fake_document,chunk_size=chunk_size)
     # print("CHUNKS: ", chunks)
 
     doc_content:list[str] = []
 
     # Assert

     assert len(chunks) == expected_chunks
     
     for i,chk in enumerate(chunks):
         
         assert len(chk.content.split()) <= chunk_size
         assert chk.document_id == fake_document.id
         assert chk.tenant_id == fake_document.tenant_id
         assert chk.allowed_groups == fake_document.allowed_groups
 
         doc_content.append(chk.content)
 
         assert chk.chunk_index == i
         assert chk.id == f"{fake_document.id}:chunk-{i}"
         
 
     reconstructed_content = " ".join(doc_content)
 
     assert reconstructed_content.split() == fake_document.content.split()


def test_chunk_empty_content():

    fake_document = Document(
            title="One Piece Live Action",
            id="doc-1",
            content="",
            tenant_id="Bank-Kara",
            allowed_groups=["risk", "executive"]
            
        )
        
    chunk_size = 3

    chunks = ChunkingService().chunk_document(document=fake_document, chunk_size=chunk_size)

    assert chunks == []

@pytest.mark.parametrize(
    "chunk_size,expected_chunks",
    [
        (0,[]),
        (-2,[])
    ]
)
def test_zero_negative_chunk_size(chunk_size,expected_chunks):

    fake_document = Document(
            title="One Piece Live Action",
            id="doc-1",
            content="The One Piece Live Action was made with effort and heart",
            tenant_id="Bank-Kara",
            allowed_groups=["risk", "executive"]
            
        )

    chunks = ChunkingService().chunk_document(document=fake_document, chunk_size=chunk_size)

    assert chunks == expected_chunks

