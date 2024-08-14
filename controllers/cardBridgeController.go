package controllers

import (
	"card-bridge/database"
	"card-bridge/entities"
	"card-bridge/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCards(c *gin.Context) {
	var result gin.H

	card, err := repositories.GetCards(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": card,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCard(c *gin.Context) {
	var card entities.Card
	idCard := c.Param("id")

	// err := c.BindJSON(&card)

	// if err != nil {
	// 	panic(err)
	// } handle menggunakan panic

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// } skip JSON Binding, karena menggunakan praram dari url

	card.Id = idCard

	err := repositories.InsertCard(database.DbConnection, card)

	// if err != nil {
	// 	panic(err)
	// } handle menggunakan panic

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, card)
}

func DeleteCard(c *gin.Context) {
	var card entities.Card
	idCard := c.Param("id")

	card.Id = idCard

	err := repositories.DeleteCard(database.DbConnection, card)
	if err != nil {
		// panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus", "id": idCard})
}
