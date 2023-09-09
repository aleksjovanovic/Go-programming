package main

import (
	"github.com/aleksjovanovic/go-crud/configs"
	"github.com/aleksjovanovic/go-crud/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Team{})
}
