package playlist

import (
	"libify/song"
	"time"
)

type Playlist struct {
	ID        uint        `gorm:"primarykey" json:"id"`
	Name      string      `json:"name"`
	Songs     []song.Song `json:"songs" gorm:"many2many:playlist_songs;"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type Filter struct {
	ID   uint
	Name string
}
