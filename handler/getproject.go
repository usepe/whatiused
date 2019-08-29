package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/usepe/whatisused/domain"
)

// GetProject get a single project by id
func GetProject(findProjectByID func(string) *domain.ProjectEnvelope) func(c *gin.Context) {
	return func(c *gin.Context) {
		var projectID = c.Param("id")

		c.JSON(200, findProjectByID(projectID))
	}
}
