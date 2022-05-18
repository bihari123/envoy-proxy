package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcom to app 3!")
	})

	r.GET("/app2", func(c *gin.Context) {
		c.String(200, "split routed to app 3!")
	})



	r.Run(":3333")
}
