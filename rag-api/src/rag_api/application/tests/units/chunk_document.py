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
    overlap = 2

    # Act
    chunks = ChunkingService().chunk_document(
        document=fake_document,
        chunk_size=chunk_size,
        overlap=overlap
    )
    print("CHUNKS: ", chunks)

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


@pytest.mark.parametrize(
    "content, chunk_size,overlap, expected_chunks",
    [
        ("The One Piece", 5, 1, 1),    # content < chunk_size
        ("The One Piece Live Action is the Best", 8, 1, 1),    # content == chunk_size
        ("The One Piece", 2, 0, 2),    # content > chunk_size
        ("The One Piece", 1, 0, 3),     # 1 word == chunk_size
        ("The One Piece Live Action was made with heart and effort.", 4, 3, 8)
    ]
)
def test_chunk_sizes(content, chunk_size, expected_chunks, overlap):

     # Arrange
     fake_document = Document(
         title="One Piece Live Action",
         id="doc-1",
         content=content,
         tenant_id="Bank-Kara",
         allowed_groups=["risk", "executive"]
         
     )
      
     # Act
     chunks = ChunkingService().chunk_document(document=fake_document,chunk_size=chunk_size,overlap=overlap)
     print("CHUNKS: ", chunks)
 
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
         
 
     # reconstructed_content = " ".join(doc_content)
 
     # assert reconstructed_content.split() == fake_document.content.split()


def test_chunk_empty_content():

    fake_document = Document(
            title="One Piece Live Action",
            id="doc-1",
            content="",
            tenant_id="Bank-Kara",
            allowed_groups=["risk", "executive"]
            
        )
        
    chunk_size = 3
    overlap = 1

    chunks = ChunkingService().chunk_document(document=fake_document, chunk_size=chunk_size,overlap=overlap)

    assert chunks == []

@pytest.mark.parametrize(
    "chunk_size, overlap, expected_error",
    [
        (0,2, "invalid chunk_size"),
        (2,3, "invalid overlap value"),
        (-2,1, "invalid chunk size" ),
        (3,-1, "invalid overlap value"),
        (3,3, "invalid overlap equal chunk size")
    ]
)
def test_zero_negative_chunk_size(chunk_size,overlap,expected_error):

    fake_document = Document(
            title="One Piece Live Action",
            id="doc-1",
            content="The One Piece Live Action was made with effort and heart",
            tenant_id="Bank-Kara",
            allowed_groups=["risk", "executive"]
            
        )
        
    with pytest.raises(ValueError) as invalid_entries:
    
        _ = ChunkingService().chunk_document(
            document=fake_document, 
            chunk_size=chunk_size,
            overlap=overlap
        )

        assert invalid_entries.match == expected_error
    

@pytest.mark.parametrize(
    "content, chunk_size, overlap",
    [
        ("One Two Three Four Five Six Seven", 3, 1 ),
        ("One Two Three Four Five Six Seven", 3, 2 ),
        ("One Two Three Four Five Six Seven", 2, 1 ),
        ("One Two Three Four Five Six Seven", 4, 1 ),
        ("One Two Three Four Five Six Seven", 4, 2 ),
        ("One Two Three Four Five Six Seven", 4, 3 ),
    ]
)
def test_overlap_chunk_content(content, chunk_size, overlap):

    # Arrange
    fake_document = Document(
        title="One Piece Live Action",
        id="doc-1",
        content=content,
        tenant_id="Bank-Kara",
        allowed_groups=["risk", "executive"]
    )
    

    # Act
    chunks = ChunkingService().chunk_document(
        document = fake_document,
        chunk_size = chunk_size,
        overlap = overlap
    )

    print("CHUNKS: ", chunks)

    # content_doc:list[str] = []

    # Assert
    for i in range(1,len(chunks)):

      previous_words = chunks[i-1].content.split()
      current_words = chunks[i].content.split()

      # The last overlap words of previous slice equal
      # the first overlap words of the current slice  
      assert previous_words[-overlap:] == current_words[:overlap] 

      