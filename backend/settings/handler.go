package settings

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Get(c *gin.Context) {
	var settings Settings
	tx := h.DB.First(&settings)
	if tx.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if tx.Error != nil {
		logrus.Errorf("error occurred while fetching settings: %s", tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, settings)
}

func (h *Handler) Update(c *gin.Context) {
	var settings Settings
	err := c.BindJSON(&settings)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = settings.HashPassword()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tx := h.DB.Model(Settings{}).Where("id = ?", 1).Updates(settings)
	if tx.Error != nil {
		logrus.Errorf("Failed to save settings: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, settings)
}