from src.rag_api.application.services import ChunkingService
from src.rag_api.domain.models import Document

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

def test_negative_chunk_size():

    fake_document = Document(
            title="One Piece Live Action",
            id="doc-1",
            content="",
            tenant_id="Bank-Kara",
            allowed_groups=["risk", "executive"]
            
        )
        
    chunk_size = -2

    chunks = ChunkingService().chunk_document(document=fake_document, chunk_size=chunk_size)

    assert chunks == []

def test_exact_chunk_size_len_content():

    fake_document = Document(
            title="One Piece Live Action",
            id="doc-1",
            content="The One Piece",
            tenant_id="Bank-Kara",
            allowed_groups=["risk", "executive"]
            
        )
        
    chunk_size = 3

    chunks = ChunkingService().chunk_document(document=fake_document, chunk_size=chunk_size)

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

def test_smaller_content_chunk_size():

    fake_document = Document(
            title="One Piece Live Action",
            id="doc-1",
            content="The One Piece",
            tenant_id="Bank-Kara",
            allowed_groups=["risk", "executive"]
            
        )
        
    chunk_size = 5

    chunks = ChunkingService().chunk_document(document=fake_document, chunk_size=chunk_size)

    content:list[str] = []

    # Assert
    for i,chk in enumerate(chunks):
        
        assert len(chk.content.split()) <= chunk_size
        assert chk.document_id == fake_document.id
        assert chk.tenant_id == fake_document.tenant_id
        assert chk.allowed_groups == fake_document.allowed_groups

        content.append(chk.content)

        assert chk.chunk_index == i
        assert chk.id == f"{fake_document.id}:chunk-{i}"
        

    reconstructed_content = " ".join(content)

    assert reconstructed_content.split() == fake_document.content.split()