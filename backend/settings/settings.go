package settings

import "golang.org/x/crypto/bcrypt"

type Settings struct {
	ID             uint   `gorm:"primaryKey"`
	SongsDirectory string `json:"songsDirectory"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"-"`
}

func (s *Settings) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	s.Password = string(bytes)
	return err
}
