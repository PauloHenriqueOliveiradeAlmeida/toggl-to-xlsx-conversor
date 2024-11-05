package togglResponse

type WorkspaceResponse struct {
	Id             int    `json:"id"`
	OrganizationId int    `json:"organization_id"`
	Name           string `json:"name"`
}
