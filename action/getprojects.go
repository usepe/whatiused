package action

import "github.com/usepe/whatisused/domain"

// AGetProjects Type Export
type AGetProjects func() []*domain.Project

// GetProjects Return all projects
func GetProjects(projects domain.Projects) func() []*domain.Project {
	return func() []*domain.Project {
		return projects.All()
	}
}
