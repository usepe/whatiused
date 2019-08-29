package infrastructure

import "github.com/usepe/whatisused/domain"

var baseList = []*domain.Project{
	domain.NewProject("1", "a", "kotlin"),
	domain.NewProject("2", "b", "kotlin"),
	domain.NewProject("3", "c", "java"),
	domain.NewProject("4", "d", "kotlin"),
	domain.NewProject("5", "e", "kotlin"),
	domain.NewProject("6", "f", "kotlin"),
	domain.NewProject("7", "g", "kotlin"),
	domain.NewProject("8", "h", "kotlin"),
	domain.NewProject("9", "i", "golang"),
	domain.NewProject("10", "j", "javascript"),
	domain.NewProject("11", "k", "kotlin"),
	domain.NewProject("12", "m", "kotlin"),
	domain.NewProject("13", "n", "kotlin"),
	domain.NewProject("14", "o", "kotlin"),
	domain.NewProject("15", "p", "kotlin"),
	domain.NewProject("16", "q", "kotlin"),
	domain.NewProject("17", "r", "kotlin"),
	domain.NewProject("18", "s", "kotlin"),
	domain.NewProject("19", "t", "kotlin"),
	domain.NewProject("20", "u", "kotlin"),
}

// InMemoryProjects contains a memory based repository of projects
type InMemoryProjects struct {
	list []*domain.Project
}

// NewInMemoryProjects returns an instance of an InMemoryProjects
func NewInMemoryProjects() *InMemoryProjects {
	return &InMemoryProjects{list: baseList}
}

// All returns a list of projects for this in memory repository
func (o *InMemoryProjects) All() []*domain.Project {
	return o.list
}

// GetByID returns the project given by ID from this in memory repository
func (o *InMemoryProjects) GetByID(ID string) *domain.Project {
	var project *domain.Project

	for _, v := range o.list {
		if v.ID == ID {
			project = v
		}
	}

	return project
}
