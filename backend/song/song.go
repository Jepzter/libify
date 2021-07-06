package song

import "time"

type Song struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Title        string    `json:"title"`
	Artist       string    `json:"artist"`
	Duration     string    `json:"duration"`
	FileLocation string    `json:"fileLocation" gorm:"index:file_location_idx,unique"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Filter struct {
	ID     uint
	Title  string
	Artist string
}
