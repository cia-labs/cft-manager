package app

import (
	"github.com/gin-gonic/gin"
)

func RunApp() {

	r := gin.Default()

	// API versioning
	v1 := r.Group("/api/v1")
	{
		v1.GET("/tunnels", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "tunnels",
			})
		})

	}

	// Run the server
	r.Run(":8080")
}
