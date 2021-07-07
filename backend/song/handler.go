package song

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"libify/settings"
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

	var songs []Song
	tx := h.DB.Find(&songs, "title like ?", filter.Title+"%")
	if tx.Error != nil {
		logrus.Errorf("error occurred while filtering songs with filters %+v:\n%s", filter, tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, songs)
}

func (h *Handler) Get(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var song Song
	tx := h.DB.First(&song, ID)
	if tx.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if tx.Error != nil {
		logrus.Errorf("error occurred while fetching song %d: %s", ID, tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, song)
}

func (h *Handler) Create(c *gin.Context) {
	var song Song
	err := c.Bind(&song)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(song.Title) == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	created, err := h.CreateSong(song)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(201, created)
}

func (h *Handler) CreateSong(song Song) (Song, error) {
	createdAt := time.Now()
	song.UpdatedAt = createdAt
	song.CreatedAt = createdAt
	tx := h.DB.Create(&song)
	if tx.Error != nil {
		logrus.Warnf("error occurred while creating song %+v:\n%s", song, tx.Error)
		return Song{}, tx.Error
	}
	var createdSong Song
	h.DB.First(&createdSong, song.ID)
	return createdSong, nil
}

func (h *Handler) Sync(c *gin.Context) {
	var setting settings.Settings
	tx := h.DB.First(&setting)
	if tx.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	sync := Synchronize{
		DB:      h.DB,
		Setting: setting,
	}
	songs := sync.Sync()
	for _, song := range songs {
		song, err := h.CreateSong(song)
		if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
			logrus.Warnf("failed to import song %+v", song)
		}
		if song.ID == 0 {
			continue
		}
		logrus.Info("Imported song ", song)
	}
	logrus.Info("song sync complete")
}
