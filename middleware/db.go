package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rexposadas/locationhistory/models"
)

// SetContext set the DB.
func SetContext(db *models.OrderDB, expiry int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)

		c.Set("EXPIRY", expiry)
		c.Next()
	}
}
