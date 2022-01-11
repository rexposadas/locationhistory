package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rexposadas/locationhistory/models"
)

func Delete(c *gin.Context) {

	db := c.MustGet("DB").(*models.OrderDB)
	orderID := c.Param("order_id")
	if orderID == "" {
		c.JSON(500, gin.H{"error": "order id is required"})
		return
	}

	// TODO: return a message if the order id doesn't exits.
	// For now, that's not a requirement.
	db.Delete(orderID)

	c.JSON(200, nil)
}
