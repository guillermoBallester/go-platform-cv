package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) HandleHome(c *gin.Context) {
	ctx := c.Request.Context()

	skills, err := r.cvSvc.GetSkills(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	experiences, err := r.cvSvc.GetExperiences(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":       "Mi CvService - Platform Engineer",
		"Skills":      skills,
		"Experiences": experiences,
	})
}
