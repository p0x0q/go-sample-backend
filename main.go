package main

import (
	"github.com/gin-gonic/gin"

	// docsのディレクトリを指定
	_ "aipco-backend/docs" // ←追記

	ginSwagger "github.com/swaggo/gin-swagger" // ←追記
)

func main() {
	r := gin.Default()

	// 下記を追記することで`http://localhost:8080/swagger/index.html`を叩くことでswagger uiを開くことができる
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
