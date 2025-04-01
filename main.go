package main

import (
	"fmt"
	"gee_gin/utils"
)

func main() {
	// fmt.Println("Gee_Gin")
	// r := gee.New()
	// r.GET("/hello", func(c *gee.Context) {
	// 	fmt.Println("123456")
	// 	c.JSON(200, map[string]string{
	// 		"data": "-",
	// 		"code": "200",
	// 		"msg":  "testtesttest",
	// 	})
	// })

	// r.GET("/doing", func(c *gee.Context) {
	// 	fmt.Println("654321")
	// 	c.HTML(200, "<h1>Do it</h1>")
	// })
	// r.Run(":8080")
	tree := utils.NewPatternTrieRoot()
	tree.Insert("/baidu/video")
	tree.Insert("/baidu/music")
	tree.Insert("/user")
	tree.Insert("/user/salary")
	fmt.Println(tree.Search("/user/salary"), tree)
}
