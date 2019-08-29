package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/usepe/whatisused/domain"
)

// ListProjects Handler for list projects
func ListProjects(getProjects func() []*domain.Project) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, domain.NewProjectList(getProjects()))
	}
}
