package domain

// Project struct
type Project struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
}

// NewProject Make a new project given a name and language
func NewProject(id, name, language string) *Project {
	return &Project{ID: id, Name: name, Language: language}
}
