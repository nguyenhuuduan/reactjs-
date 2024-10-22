package routes

import (
	"dating_app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Routes đăng ký người dùng
	router.POST("/register", controllers.RegisterUser)

	// Routes cho tính năng quên mật khẩu
	router.POST("/forgot-password", controllers.ForgotPassword)
	router.POST("/verify-otp", controllers.VerifyOTP)
	router.POST("/reset-password", controllers.ResetPassword)

	// Routes cho tính năng chỉnh sửa thông tin người dùng bằng ID
	router.GET("/user/:id", controllers.GetUserProfile)    // Tìm kiếm bằng ID
	router.PUT("/user/:id", controllers.UpdateUserProfile) // Cập nhật bằng ID
}
