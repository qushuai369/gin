package user

import (
	"database/sql"
	"fmt"
	"gin/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"net/http"
	"os"
	"path"
	"strconv"
	"time"
	"math/rand"
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

func PostData(c *gin.Context)  {
	message := c.PostFormArray("message")
	fmt.Printf("类型：%T，值：%v\n",message,message)
}

//上传
func PostFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err ==nil {
		Path := "./public/image/"
		t := time.Now()
		date := t.Format("20060102")
		pathTmp := Path + "/ " + date +"/"
		if isDirExists(pathTmp) {
			fmt.Println("目录存在")
		} else {
			fmt.Println("目录不存在")
			err := os.Mkdir(pathTmp, 0777)
			if err != nil {
				//log.Fatal(err)
				c.JSON(200, gin.H{"status": -1, "msg": "创建目录失败",})
			}
		}
		//文件名
		file_name := strconv.FormatInt(time.Now().Unix(),10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + path.Ext(file.Filename)
		uperr := c.SaveUploadedFile(file, pathTmp + file_name)
		if(uperr==nil) {
			c.JSON(200, gin.H{"status": 1, "msg": "上传成功","data":date+"/"+file_name})
		} else {
			c.JSON(200, gin.H{"status": -2, "msg": "上传失败",})
		}
		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	} else {
		c.JSON(200, gin.H{"status": -1, "msg": "上传失败",})
	}
}

func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)//err!=nil,使用os.IsExist()判断为false,说明文件或文件夹不存在
	} else {
		return fi.IsDir()//err==nil,说明文件或文件夹存在
	}
	panic("not reached")
}
