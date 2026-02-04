package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) HandleHome(c *gin.Context) {
	skills, err := r.cvSvc.GetSkills(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":  "Mi CV - Platform Engineer",
		"Skills": skills,
	})
}
