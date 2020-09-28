package main

import (
	"fmt"
	"gin/app/user"
	"github.com/gin-gonic/gin"
	"gin/common"
)

func main() {
	fmt.Println(common.Date())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/someGet", getting)
	r.GET("/user/get", user.GetAllUser)
	r.GET("/user/getUsers", user.GetUsers)
	r.GET("/home/index", user.Index)
	r.GET("/home/tt", user.Tt)
	r.POST("/home/Postdata", user.PostData)
	r.POST("/home/Postfile", user.PostFile)
	r.Run(":8181") // listen and serve on 0.0.0.0:8080

}

func getting(c *gin.Context) {
	var aa = [...]int{1, 2, 3}
	var bb = []int{1, 2, 3}
	//aa[3] = 4
	//bb[3] = 4
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Printf("%v类型：%T\n", aa, aa)
	fmt.Printf("类型：%T\n", bb)
	c.JSON(200, gin.H{
		"message": "someGet",
	})
}
