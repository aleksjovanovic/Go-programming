package models

import (
	// "time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID              uint
	TEAM            string
	MATCHES_PLAYED  int
	WIN             int
	DRAW            int
	LOSE            int
	GOAL_DIFFERENCE int
	POINTS          uint
	FORM            string
	// CREATED_AT      time.Time
	// UPDATED_AT      time.Time
	// DELETED_AT      time.Time
}

func (Team) TableName() string {
	return "teams"
}
