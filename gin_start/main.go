package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)


type Person struct{
	ID int `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

type LoginForm struct {
	User string	`json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpForm struct {
	Age uint8 `json:"age" binding:"gte=1,lte=130"`
	Name string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"` 
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context){
		t := time.Now()
		c.Set("example", "123456")
		c.Next()

		end := time.Since(t)
		fmt.Println(end)
		status := c.Writer.Status()
		fmt.Println(status)
	}
}

func main() {
	r := gin.Default()
	r.Use(MyLogger())
	goodsGroup := r.Group("goods")
	{
		goodsGroup.GET("", goodsList)
		goodsGroup.GET("/:name/:id", goodsDetail)
		goodsGroup.POST("", createGoods)
	}
	r.GET("/welcome", welcome)
	r.POST("/formpost", formPost)
	r.POST("/post", getPost)
	r.POST("/getjson", getJson)
	r.POST("/getproto", getProto)


	r.POST("/loginJSON", func(c *gin.Context){
		var logForm LoginForm
		if err := c.ShouldBind(&logForm); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
	})

	r.POST("/signup", func(c *gin.Context){
		var signUpForm SignUpForm
		if err := c.ShouldBind(&signUpForm); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "注册成功",
		})
	})

	go func(){
		r.Run("127.0.0.1:8888") // 监听并在 0.0.0.0:8080 上启动服务
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	fmt.Println("close server")


	
}

func getProto(c *gin.Context){

}

func getJson(c *gin.Context){
	var msg struct{
		Name string `json:"user"`
		Message string
		Number int
	}

	msg.Name = "bobby"
	msg.Message = "test message"
	msg.Number = 20
	c.JSON(http.StatusOK, msg)
}
func getPost(c *gin.Context){
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name") 
	message := c.DefaultPostForm("message", "信息")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"page": page,
		"name": name,
		"message": message,
	})
}


func welcome(c *gin.Context){
	firstName := c.DefaultQuery("firstName", "bobby")
	lastName := c.DefaultQuery("lastName", "doddy")
	c.JSON(http.StatusOK, gin.H{
		"firstName": firstName,
		"lastName": lastName,
	})
}
func formPost(c *gin.Context){
	messages := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"message": messages,
		"nick": nick,
	})
}
func goodsList(context *gin.Context){
	// context.JSON(http.StatusOK, gin.H{
	// 	"name": "goodsList",
	// })
	var person Person
	if err:=context.ShouldBindUri(&person); err != nil {
		context.Status(404)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"name": person.Name,
		"id": person.ID,
	})
}

func goodsDetail(context *gin.Context){
	id := context.Param("id")
	action := context.Param("action")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
		"action": action,
	})
	
}

func createGoods(context *gin.Context){
	
}