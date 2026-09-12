
from dataclasses import dataclass


@dataclass
class User:
    user_id: str
    tenant_id: str
    groups: list[str]


@dataclass
class Document:
    title: str
    id: str
    content: str
    tenant_id: str
    allowed_groups: list[str]