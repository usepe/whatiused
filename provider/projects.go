package provider

import (
	"sync"

	"github.com/usepe/whatisused/domain"
	"github.com/usepe/whatisused/infrastructure"
)

var once sync.Once

var (
	instance domain.Projects
)

// GetProjects returns the repository set up for projects
func GetProjects() domain.Projects {
	once.Do(func() {
		instance = infrastructure.NewInMemoryProjects()
	})

	return instance
}
