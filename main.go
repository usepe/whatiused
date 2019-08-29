package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usepe/whatisused/action"
	"github.com/usepe/whatisused/handler"
	"github.com/usepe/whatisused/provider"
)

func main() {
	r := gin.Default()

	// Repository
	var projects = provider.GetProjects()

	// Actions
	var getProjects = action.GetProjects(projects)
	var findProjectByID = action.FindProjectByID(projects)

	// Routes
	r.GET("/projects", handler.ListProjects(getProjects))
	r.GET("/projects/:id", handler.GetProject(findProjectByID))

	r.Run()
}
