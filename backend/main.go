package main

import (
	"github.com/sirupsen/logrus"
	"libify/settings"
	"os"
	"strconv"
)

func main() {
	fromDockerVolume, err := strconv.ParseBool(os.Getenv("MUSIC_FROM_DOCKER_VOLUME"))
	if err != nil {
		fromDockerVolume = false
	}
	db, err := connectDB()
	if err != nil {
		logrus.Fatal("could not connect to database: ", err)
	}
	volumePath := os.Getenv("VOLUME_PATH")
	err = settings.Setup(db, fromDockerVolume, volumePath)
	if err != nil {
		logrus.Fatal("error setting up settings with docker volume")
	}
	server := WebServer{DB: db}
	err = server.Start(volumePath)
	if err != nil {
		logrus.Fatal("could not start libify: ", err)
	}
}
