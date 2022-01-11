package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rexposadas/locationhistory/models"
)

func List(c *gin.Context) {

	db := c.MustGet("DB").(*models.OrderDB)
	orderID := c.Param("order_id")
	if orderID == "" {
		c.JSON(500, gin.H{"error": "order id is required"})
		return
	}

	max, err := setMax(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	list := db.List(orderID, max)

	c.JSON(200, list)
}

func setMax(c *gin.Context) (int, error) {
	maxStr := c.Query("max")

	if maxStr == "" {
		return 0, nil
	}

	m, err := strconv.Atoi(maxStr)
	if err != nil {
		return 0, fmt.Errorf("invalid max value of %s", maxStr)
	}
	return m, nil
}
