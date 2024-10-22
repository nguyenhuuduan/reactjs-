package main

import (
	"dating_app/config"
	"dating_app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Thêm middleware CORS cho phép yêu cầu từ mọi nguồn
	r.Use(cors.Default())

	// Kết nối tới MongoDB
	config.ConnectDatabase()

	// Đăng ký các routes
	routes.RegisterRoutes(r)

	// Chạy server tại cổng 8080
	r.Run(":8080")

	r.Use(cors.Default()) // Thêm dòng này vào main.go để bật CORS
}
