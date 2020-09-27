package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"strconv"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"sync"
)

var wg sync.WaitGroup

func Index(c *gin.Context) {
	//获取当前计算机上面的 Cup 个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
	//可以自己设置使用多个 cpu
	runtime.GOMAXPROCS(1)
	//cpuNum = runtime.NumCPU()
	//fmt.Println("cpuNum=", cpuNum)
	t1 := time.Now()
	n1 := 100000000
	var r1 int
	var r2 int
	var r3 int
	var r4 int
	var r5 int
	var r6 int
	var r7 int
	var r8 int
	wg.Add(8)
	go suan(n1, &r1)
	go suan(n1, &r2)
	go suan(n1, &r3)
	go suan(n1, &r4)
	go suan(n1, &r5)
	go suan(n1, &r6)
	go suan(n1, &r7)
	go suan(n1, &r8)
	wg.Wait()
	t2 := time.Since(t1)
	fmt.Println(t2)
	fmt.Println(r1)
	//fmt.Println(r2)
	//fmt.Println(r3)
	//fmt.Println(r4)

}

func suan(n int, r *int) {
	for i := 0; i < n; i++ {
		*r = *r + i
	}
	wg.Done()
}

func Tt(c *gin.Context) {
	pathTmp := "/ " + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) +"/"
	c.JSON(200, gin.H{
		"status": 1,
		"msg":    "成功",
		"data":   pathTmp,
	})
	fmt.Println(pathTmp)
	fmt.Printf("%T",time.Now().Unix())
}
