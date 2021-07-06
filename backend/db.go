package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"libify/playlist"
	"libify/settings"
	"libify/song"
)

func connectDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "hej123456", "127.0.0.1", "3306", "libify")
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(playlist.Playlist{}, song.Song{}, settings.Settings{})
	if err != nil {
		return nil, err
	}
	return db, nil
}