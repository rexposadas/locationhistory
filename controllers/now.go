package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rexposadas/locationhistory/models"
)

// Now adds order details to our db.
// sample call and payload:
// POST /location/def456/now
// {
// 	"lat": 12.34,
// 	"lng": 56.78
// }
func Now(c *gin.Context) {
	db := c.MustGet("DB").(*models.OrderDB)
	expiry := c.MustGet("EXPIRY").(int)

	orderID := c.Param("order_id")

	if orderID == "" {
		c.JSON(500, gin.H{"error": "order id is required"})
		return
	}

	loc := models.Location{}
	if err := c.Bind(&loc); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	db.Add(orderID, loc, expiry)

	c.JSON(200, nil)
}
