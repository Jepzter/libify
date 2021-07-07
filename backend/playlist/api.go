package playlist

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *API) Register() {
	playlistV1 := a.Router.Group("/api/libify/playlist-v1")
	{
		handler := Handler{DB: a.DB}
		playlistV1.GET("", handler.Find)
		playlistV1.GET("/:id", handler.Get)
		playlistV1.POST("", handler.Create)
		playlistV1.PUT("", handler.Update)
	}
}
