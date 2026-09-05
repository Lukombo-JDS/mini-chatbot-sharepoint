package domain

type Document struct{
    ID string
    Name string
    Content string
    TenantID string
    AllowGroups  []string
}