package song

import (
	"github.com/dhowden/tag"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"libify/settings"
	"os"
	"path/filepath"
	"strings"
)

type Synchronize struct {
	DB      *gorm.DB
	Setting settings.Settings
}

func (s *Synchronize) Sync() []Song {
	if strings.TrimSpace(s.Setting.SongsDirectory) == "" {
		return []Song{}
	}
	if strings.TrimSpace(s.Setting.Host) == "" {
		return s.fromVolume()
	}
	return s.fromRemote()
}

func (s *Synchronize) fromVolume() []Song {
	var songs []Song
	err := filepath.Walk(s.Setting.SongsDirectory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".mp3") {
			f, err := os.Open(path)
			defer f.Close()
			if err != nil {
				logrus.Warnf("could not import file %s: %v", path, err)
				return nil
			}
			mp3, err := tag.ReadFrom(f)
			if err != nil {
				logrus.Warnf("could not import file %s: %v", path, err)
				return nil
			}
			songs = append(songs, Song{
				Title:        mp3.Title(),
				Artist:       mp3.Artist(),
				Duration:     "N/A",
				FileLocation: strings.ReplaceAll(path, s.Setting.SongsDirectory, ""),
			})
		}
		return nil
	})
	if err != nil {
		logrus.Warn("problem synchronizing songs: ", err)
	}
	return songs
}

func (s *Synchronize) fromRemote() []Song {
	return []Song{}
}
