import httpx
from src.rag_api.domain.models import Document, User
from src.rag_api.infrastructure.sharepoint_client import SharePointClient
import pytest

def handler_status_ok(request: httpx.Request):

    document = {
        "ID": "docs-1",
        "Title": "One Piece",
        "Content": "Emperor Crew",
        "TenantID": "Bank-Kara",
        "AllowedGroups": ["risk","executive"]
    }

    assert request.method == "GET"
    assert request.url.path == "/documents"
    assert request.headers.get("X-User-ID") == "Nami"
    assert request.headers.get("X-Tenant-ID") == "Bank-Kara"
    assert request.headers.get("X-Groups") == "risk,executive"
    
    return httpx.Response(
        httpx.codes.OK,
        json={
            "status": "ok",
            "documents": [document]
        }
    )

def handler_status_unauthorized(request: httpx.Request):

    assert request.method == "GET"
    assert request.url.path == "/documents"

    
    return httpx.Response(
        httpx.codes.UNAUTHORIZED,
        json={
            "error": "missing user identity"
        }
    )

def test_sharepoint_client():

    ######## ARRANGE ######

    user = User(
        user_id="Nami",
        tenant_id="Bank-Kara",
        groups=["risk","executive"]
    )

    transport = httpx.MockTransport(handler_status_ok) # fake server with fake handler

    client = httpx.Client(transport=transport) # fake client

    ######### ACT #######

    # testing SharePoint client methods
    docs = SharePointClient(client).list_documents(user)

    ######## ASSERT #######

    assert len(docs) == 1 # right amount of documents
    assert isinstance(docs,list) # right type of response
    for d in docs:
        assert isinstance(d, Document)
    #right docs IDs
    assert docs[0].id == "docs-1"  
    assert docs[0].tenant_id == "Bank-Kara" 
    assert docs[0].allowed_groups == ["risk", "executive"] # right allowed groups    

def test_sharepoint_client_error():

    user = User(
        user_id="Baggy",
        tenant_id="Bank-Cross-Guild",
        groups=["risk", "executive"]
    )

    transport = httpx.MockTransport(handler=handler_status_unauthorized)

    client = httpx.Client(transport=transport)

    with pytest.raises(httpx.HTTPStatusError, match="401 Unauthorized"):
        
        _ = SharePointClient(client).list_documents(user)

