package domain

// ProjectList struct
type ProjectList struct {
	Projects []*Project `json:"projects"`
}

// NewProjectList Make a new project list with the given projects
func NewProjectList(list []*Project) *ProjectList {
	return &ProjectList{Projects: list}
}
