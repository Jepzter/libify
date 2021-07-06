package settings

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *API) Register() {
	settingsV1 := a.Router.Group("/api/libify/settings-v1")
	{
		handler := Handler{DB: a.DB}
		settingsV1.GET("", handler.Get)
		settingsV1.PUT("", handler.Update)
	}
}
