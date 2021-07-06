package song

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *API) Register() {
	songV1 := a.Router.Group("/api/libify/song-v1")
	{
		handler := Handler{DB: a.DB}
		songV1.GET("", handler.Find)
		songV1.GET("/:id", handler.Get)
		songV1.POST("", handler.Create)
		songV1.PUT("/sync", handler.Sync)
	}
}
