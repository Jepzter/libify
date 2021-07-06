package settings

import (
	"gorm.io/gorm"
	"strings"
)

func Setup(db *gorm.DB, fromDockerVolume bool, volumePath string) error {
	if !fromDockerVolume || strings.TrimSpace(volumePath) == "" {
		return nil
	}
	var settings Settings
	tx := db.FirstOrCreate(&settings, Settings{
		SongsDirectory: volumePath,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
