package domain

// ProjectEnvelope struct
type ProjectEnvelope struct {
	Project *Project `json:"project"`
}

// NewProjectEnvelope Make a new project envelope (a project for json response)
func NewProjectEnvelope(project *Project) *ProjectEnvelope {
	return &ProjectEnvelope{Project: project}
}
