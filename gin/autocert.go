//lets encrypt
//很久很久以前，想要一张ssl证书就必须每年为一个域名花很多钱。而现在，Let's Encrypt改变了这一切。
//这是一个非营利性机构，它提供免费证书，并且提供HTTP API来获得证书。API允许自动化处理这一过程。
//在Let’s Encrypt出现之前，你可能会购买一个证书，而这仅仅是一串字节而已。你会把证书存放在一个文件中，并配置你的网络服务器来使用它。
//有了Let’s Encrypt以后，你就能够使用他们的API来免费获得证书了，而且这一过程是在你的服务器启动后自动完成的。
//值得庆幸的是，所有与API进行对话的艰苦工作都已由其他人做完了。我们只需要安装插件就可以了。
//有两个Go库可以实现Let’s Encrypt支持。

package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
)

func automain() {
	r := gin.Default()
	r.GET("/ping")
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("xxx", "xxx"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
}
