package http

import (
	"github.com/gin-gonic/gin"
	"github.com/guillermoBallester/go-platform-cv/internal/service"
	"net/http"
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

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(w, req)
}
