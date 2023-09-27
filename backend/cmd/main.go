package main

import (
	"backend/api"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/db/mongo"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	c := config.GetConfig()

	opts := db.NewDatabaseOptions(c.MongoConfig.ConnectionURI, c.MongoConfig.Database, time.Millisecond*time.Duration(c.MongoConfig.Timeout))
	db := mongo.NewMongoBackend(opts)

	err := db.Connect()
	if err != nil {
		logrus.Panic(err)
	}

	s := api.NewServer(db)

	if err := s.Serve(&c); err != nil {
		logrus.Panic(err)
	}
}