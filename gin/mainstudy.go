// httprouter
// go build -tags=jsoniter
package main

import (
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//默认已经连接了logger and recovery 中间件
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//匹配/user/john/ 或者 /user/john/send
	//无路由匹配/user/john，则会重定向到/user/john/
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})

	// /welcome?firstname=1&lastname=2
	r.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "guest")
		lastname := context.Query("lastname")
		message := firstname + lastname
		context.String(http.StatusOK, message)
	})

	r.POST("/user/:name/*action", func(context *gin.Context) {
		context.FullPath()
	})
	r.POST("/post", func(context *gin.Context) {
		//表单
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
		//query+post表单/post?id=1234&page=1&ids[a]=1234 表单name=manu names[first]=1
		//map作为查询字段或post表单
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")

		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")
		fmt.Println(id, page, name, ids, names)

	})
	r.MaxMultipartMemory = 8 << 20 //8mib //限制上传大小（默认32mib）
	dst := "/c/xiyou"
	r.POST("upload", func(context *gin.Context) {
		//单文件
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		context.SaveUploadedFile(file, dst) //上传到指定路径
		//多文件
		from, _ := context.MultipartForm()
		files := from.File["upload[]"]
		for _, val := range files {
			log.Println(val.Filename)
			context.SaveUploadedFile(file, dst)
		}

	})

	//路由分组
	v1 := r.Group("v1")
	{
		v1.POST("login")
		v1.POST("submit")
	}

	//默认没有中间间的空白gin
	rm := gin.New()
	//使用中间件
	//全局中间件
	//Logger 将日志写到 gin.defaulwriter 即使你设置gin_mode=release
	//默认设置gin.defaulwriter=os.stout
	rm.Use(gin.Logger())
	//Recovery 从任何panic恢复 如果出现panic 会写一个500错误
	rm.Use(gin.Recovery())

	//授权组
	//使用自定义创建的AuthRequired()中间件
	//authorized:=r1.Group("/",AuthRequired())
	//{
	//	authorized.POST("login")
	//}
	//BasicAuth 基本鉴权
	//authorized:=rm.Group("/admin",gin.BasicAuth(gin.Accounts{
	//	"foo":"bar",
	//})

	//如何写入日志文件？
	//禁用控制台颜色
	gin.DisableConsoleColor()
	//写入日志文件
	f, _ := os.Create("gin.log")
	//即写入日志文件，又在控制台显示
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//自定义日志格式
	rm.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s", params.ClientIP)
	}))
	//记录日志颜色
	gin.ForceConsoleColor()

	//模型绑定和验证（binding）
	//1,ShouldBind 使用这一个
	// 1.1 ShouldBind 可以不指定具体类型，会根据content-type判读使用了那种绑定器;绑定get/post参数;绑定html复选框
	// 1.2 ShouldBindQuery 只绑定查询字符串
	// 1.3 ShouldBindUri 绑定url "/:name/:id"
	// 1.4 ShouldBindHeader 绑定header curl -H "name:123" ...
	//2,MustBind

	r.POST("loginjson", func(context *gin.Context) {
		var login Login
		var msg ret
		msg.Name = login.User

		if err := context.ShouldBind(&login); err != nil {
			msg.Message = err.Error()
			context.JSON(http.StatusBadRequest, msg)
		}
		if login.User != "xiyou" || login.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

		//返回值：可以是xml,json,yaml,protobuf
		//gin.H=> 可以自定义
		//context.SecureJSON()//可以防止json劫持，如果返回的是数组会默认在前面加上 "while(1)"
		//context.JSONP()//在不同域中从一个服务器请求数据，如果请求参数中存在callback，添加callback到response.body
		//eg:curl http://xxxx/JSONP?callback=x 将会输出 x({xx:1,xxx:2})

		//从文件中提供数据
		context.File("xxx/xxx.go")
		//从reader提供数据
		//context.DataFromReader()

		//html渲染
		//index.html
		//<html><h1>{{.title}}</h1></html>
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "test web",
		})

		//context.Redirect()//重定向
		//发出路由重定向
		context.Request.URL.Path = "xxx"
		r.HandleContext(context)

		//!!!!!!!!中间件中使用goroutine 不能直接使用上下文，必须使用只读副本
		ccp := context.Copy()
		go func() {
			log.Println("do!" + ccp.Request.URL.Path)
		}()
	})
	//静态文件
	r.Static("/assets", "./assets")

	//自定义模版渲染器
	html := template.Must(template.ParseFiles("file1", "file2"))
	r.SetHTMLTemplate(html)

	//自定义分隔符
	r.Delims("{[{", "}]}")
	r.Run(":3000")

	//http.ListenAndServe(r)
	//自定义http配置
	s := &http.Server{
		Addr:        ":8080",
		Handler:     r,
		ReadTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

type Login struct {
	User     string `from:"user" json:"user" binding:"required"`
	Password string `from:"password" json:"password" binding:"required"`
}
type ret struct {
	Message string `json:"msg"`
	Name    string `json:"name"`
}

//gin.H 是 map[string]interface{}的快捷写法
//gin.accounts 是 map[string]string的快捷写法

//lets encrypt
//很久很久以前，想要一张ssl证书就必须每年为一个域名花很多钱。而现在，Let's Encrypt改变了这一切。
//这是一个非营利性机构，它提供免费证书，并且提供HTTP API来获得证书。API允许自动化处理这一过程。
//在Let’s Encrypt出现之前，你可能会购买一个证书，而这仅仅是一串字节而已。你会把证书存放在一个文件中，并配置你的网络服务器来使用它。
//有了Let’s Encrypt以后，你就能够使用他们的API来免费获得证书了，而且这一过程是在你的服务器启动后自动完成的。
//值得庆幸的是，所有与API进行对话的艰苦工作都已由其他人做完了。我们只需要安装插件就可以了。
//有两个Go库可以实现Let’s Encrypt支持。

func mainnew() {
	r := gin.Default()
	r.GET("/ping")
	log.Fatal(autotls.Run(r, "xxx", "xxx"))
}
