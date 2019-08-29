package action

import "github.com/usepe/whatisused/domain"

// FindProjectByID returns the project given by id
func FindProjectByID(projects domain.Projects) func(id string) *domain.ProjectEnvelope {
	return func(id string) *domain.ProjectEnvelope {
		return domain.NewProjectEnvelope(projects.GetByID(id))
	}
}
