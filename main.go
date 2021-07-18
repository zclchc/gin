package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {

		var err error
		// 读取body 数据
		//bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		//if err != nil {
		//	log.Println("err: ", err)
		//}
		//
		//// ioutil.ReadAll 读取后，body 值为空
		//// 把值放回body中
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//PostForm := c.PostForm("name")

		type name struct {
			Id int `form:"id" binding:"required,gt=10"`
		}
		//var BindJSONRes interface{}
		//BindJSONRes :=  make([]name, 0)
		//BindJSONRes :=  new(name)
		//err = c.BindJSON(&BindJSONRes)
		//if err != nil {
		//	fmt.Println("BindJSONRes", err)
		//return
		//}
		//c.JSON(http.StatusOK, BindJSONRes)

		ShouldBindres := new(name)
		err = c.ShouldBind(&ShouldBindres)
		if err != nil {
			fmt.Println("BindJSONRes", err)
			c.JSON(http.StatusAccepted, err)
			return
		}
		c.JSON(http.StatusOK, ShouldBindres)


		//BindRes := make([]name, 0)
		//err = c.Bind(&BindRes)
		//fmt.Println(err)
		//
		//fmt.Println(BindRes)
		// ShouldBindres := new(name)
		//err = c.ShouldBind(&ShouldBindres)
		//fmt.Println(err)
		//log.Println("ReadAll: ", string(bodyBytes))
		//
		////  curl -X POST http://127.0.0.1:8088/test -d 'name=123'
		//// {"PostForm":"123","bodyBytes":"name=123"}
		//c.JSON(http.StatusOK, gin.H{
		//	"bodyBytes":     string(bodyBytes),
		//	"PostForm":      PostForm,
		//	"BindJSONRes":   BindJSONRes,
		//	"BindRes":       BindRes,
		//	"ShouldBindres": ShouldBindres,
		//})
	})
	err := r.Run(":8088")
	if err != nil {
		log.Println("监听端口错误", err)
	}

}

//func createDirect(name string) {
//
//	//for i:=3; i < 8; i++ {
//	//	createDirect(strconv.Itoa(i))
//	//}
//
//	dirName := "0" +  name
//	err := os.Mkdir(dirName, os.ModePerm) //创建目录
//	//os.MkdirAll("dir1/dir2/dir3", os.ModePerm)   //创建多级目录
//	if err != nil {
//		log.Println("创建目录失败：", err)
//	}
//	for i := 1; i < 5; i++ {
//		_, err := os.Create("./" + dirName + "/0" + strconv.Itoa(i) + ".md")
//		if err != nil {
//			log.Println("创建文件失败：", err)
//		}
//	}
//
//}
