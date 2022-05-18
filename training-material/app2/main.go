package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcom to app 2!")
	})
r.GET("/app1", func(c *gin.Context) {
		c.String(200, "split routed to app2!")
	})


	r.Run(":2222")
}
