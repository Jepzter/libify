package playlist

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Find(c *gin.Context) {
	var filter Filter
	err := c.BindQuery(&filter)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var playlists []Playlist
	tx := h.DB.Where(&Playlist{
		ID: filter.ID,
		Name: filter.Name,
	}).Find(&playlists)

	if tx.Error != nil {
		logrus.Errorf("error occurred while filtering playlist with filters %+v:\n %s", filter, tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, playlists)
}

func (h *Handler) Get(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var playlist Playlist
	tx := h.DB.Preload("Songs").Find(&playlist, "id = ?", ID)
	if tx.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if tx.Error != nil {
		logrus.Errorf("error occurred while fetching playlist %d: %s", ID, tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, playlist)
}

func (h *Handler) Create(c *gin.Context) {
	var playlist Playlist
	err := c.Bind(&playlist)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(playlist.Name) == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	createdAt := time.Now()
	playlist.UpdatedAt = createdAt
	playlist.CreatedAt = createdAt
	tx := h.DB.Create(&playlist)
	if tx.Error != nil {
		logrus.Errorf("error occurred while creating playlist %+v:\n%s", playlist, tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var createdPlaylist Playlist
	h.DB.First(&createdPlaylist, playlist.ID)
	c.JSON(201, createdPlaylist)
}