package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcom to app 4!")
	})

r.GET("/app2", func(c *gin.Context) {
		c.String(200, "split route to  app 4!")
	})


	r.Run(":4444")
}
