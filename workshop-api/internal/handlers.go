package internal

import (
	"github.com/gin-gonic/gin"
)

func Handlers(r *gin.Engine, s *Server) {
	products := r.Group("products")
	products.GET("", s.FindAllProductCtrl.Handle)
	products.POST("", s.CreatProductCtrl.Handle)
}
