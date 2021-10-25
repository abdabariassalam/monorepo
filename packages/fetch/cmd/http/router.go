package http

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	Handler httpHandler
}

func NewRouter(handler httpHandler) *Router {
	return &Router{
		router:  gin.Default(),
		Handler: handler,
	}
}

func (r *Router) routes() {
	r.router.GET("/verify-token", r.Handler.VerifyToken)
	r.router.GET("/aggregate", r.Handler.Aggregator)
	r.router.GET("/currency-converter", r.Handler.CurrencyConverter)
	// set timeout

}
