package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"libify/playlist"
	"libify/settings"
	"libify/song"
)

type WebServer struct {
	DB *gorm.DB
}

func (s *WebServer) Start(streamRoot string) error {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	r.Use(cors.New(config))

	playlistAPI := playlist.API{Router: r, DB: s.DB}
	songAPI := song.API{Router: r, DB: s.DB}
	settingsAPI := settings.API{Router: r, DB: s.DB}

	playlistAPI.Register()
	songAPI.Register()
	settingsAPI.Register()

	r.Static("/api/libify/stream-v1/", streamRoot)

	return r.Run(":8085")
}