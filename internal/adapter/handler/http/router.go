package http

import (
	"github.com/gin-gonic/gin"
	"github.com/guillermoBallester/go-platform-cv/internal/service"
)

type Router struct {
	engine *gin.Engine
	cvSvc  *service.CVService
}

func NewRouter(cvSvc *service.CVService) *Router {
	g := gin.Default()

	g.Static("/assets", "./assets")

	r := &Router{
		engine: g,
		cvSvc:  cvSvc,
	}

	g.LoadHTMLGlob("templates/*.html")

	g.GET("/", r.HandleHome)

	return r
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
