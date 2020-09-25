package user

import (
	"database/sql"
	"fmt"
	"gin/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
)

//var db interface{}

type users struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUsers(c *gin.Context) {
	db, _ := sql.Open("mysql", config.Db)
	rows, _ := db.Query("SELECT name,age from users WHERE age<?", 10)
	var list []users
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			fmt.Println(err)
		}
		list = append(list, users{Name: name, Age: age})
	}
	c.JSON(200, gin.H{
		"status": 1,
		"msg":    "成功",
		"data":   list,
	})

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
		"status": 1,
		"msg":    "成功",
		"data":   aa,
	})
}
