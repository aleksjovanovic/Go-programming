package controllers

import (
	"github.com/aleksjovanovic/go-crud/configs"
	"github.com/aleksjovanovic/go-crud/models"

	"github.com/gin-gonic/gin"
)

type TeamRequestBody struct {
	ID              uint   `json:"ID"`
	TEAM            string `json:"TEAM"`
	MATCHES_PLAYED  int    `json:"MATCHES_PLAYED"`
	WIN             int    `json:"WIN"`
	DRAW            int    `json:"DRAW"`
	LOSE            int    `json:"LOSE"`
	GOAL_DIFFERENCE int    `json:"GOAL_DIFFERENCE"`
	POINTS          uint   `json:"POINTS"`
	FORM            string `json:"FORM"`
}

func TeamCreate(c *gin.Context) {

	body := TeamRequestBody{}

	c.BindJSON(&body)

	team := &models.Team{ID: body.ID, TEAM: body.TEAM, MATCHES_PLAYED: body.MATCHES_PLAYED, WIN: body.WIN, DRAW: body.DRAW, LOSE: body.LOSE, GOAL_DIFFERENCE: body.GOAL_DIFFERENCE, POINTS: body.POINTS, FORM: body.FORM}

	result := configs.DB.Create(&team)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &team)
}

func TeamGet(c *gin.Context) {
	var teams []models.Team
	configs.DB.Find(&teams)
	c.JSON(200, &teams)
	return
}

func TeamGetById(c *gin.Context) {
	id := c.Param("id")
	var team models.Team
	configs.DB.First(&team, id)
	c.JSON(200, &team)
	return
}

func TeamUpdate(c *gin.Context) {

	id := c.Param("id")
	var team models.Team
	configs.DB.First(&team, id)

	body := TeamRequestBody{}
	c.BindJSON(&body)
	data := &models.Team{ID: body.ID, TEAM: body.TEAM, MATCHES_PLAYED: body.MATCHES_PLAYED, WIN: body.WIN, DRAW: body.DRAW, LOSE: body.LOSE, GOAL_DIFFERENCE: body.GOAL_DIFFERENCE, POINTS: body.POINTS, FORM: body.FORM}

	result := configs.DB.Model(&team).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &team)
}

// soft delete

func TeamDelete(c *gin.Context) {
	id := c.Param("id")
	var team models.Team
	configs.DB.Delete(&team, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
