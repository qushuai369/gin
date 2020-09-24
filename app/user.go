package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUser() {
	fmt.Println("getUset")
}

func GetAllUser(c *gin.Context) {
	var aa = [...]int{1, 2, 3}
	var bb = []int{1, 2, 3}
	//aa[3] = 4
	//bb[3] = 4
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Printf("%v类型：%T\n", aa, aa)
	fmt.Printf("类型：%T\n", bb)
	c.JSON(200, gin.H{
		"message": "getalluser",
		"data":    aa,
	})
}
