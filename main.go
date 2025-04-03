package main

import (
	"fmt"
	gee "gee_gin/apis"
	"net/http"
)

func main() {
	fmt.Println("Gee_Gin")
	r := gee.New()
	r.GET("/hello", func(c *gee.Context) {
		fmt.Println("123456")
		c.JSON(200, map[string]string{
			"data": "-",
			"code": "200",
			"msg":  "testtesttest",
		})
	})

	// r.GET("/doing", func(c *gee.Context) {
	// 	fmt.Println("654321")
	// 	c.HTML(200, "<h1>Do it</h1>")
	// })
	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	// r.GET("/assets/*filepath", func(c *gee.Context) {
	// 	c.JSON(http.StatusOK, map[string]string{"filepath": c.Param("filepath")})
	// })
	r.Run(":8098")
}
