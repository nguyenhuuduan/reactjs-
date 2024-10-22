// controllers/user.go
package controllers

import (
	"context"
	"dating_app/config"
	"dating_app/models"
	"dating_app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Đăng ký người dùng mới
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	collection := config.GetCollection("users")
	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	user.ID = primitive.NewObjectID()
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Hàm quên mật khẩu và gửi OTP qua email
func ForgotPassword(c *gin.Context) {
	email := c.PostForm("email")

	collection := config.GetCollection("users")
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	// Kiểm tra nếu đã gửi OTP trong 30 giây qua
	currentTime := time.Now().Unix()
	timeSinceLastOtp := currentTime - user.LastOTPSentAt

	if timeSinceLastOtp < 30 {
		// Trả về thời gian còn lại trước khi có thể gửi OTP mới
		timeRemaining := 30 - timeSinceLastOtp
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error":          "Please wait before requesting a new OTP",
			"time_remaining": timeRemaining, // Trả về số giây còn lại
		})
		return
	}

	// Tạo OTP mới
	otp := utils.GenerateOTP()
	otpExpiresAt := time.Now().Add(15 * time.Minute).Unix()

	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"email": email},
		bson.M{"$set": bson.M{
			"otp":              otp,
			"otp_expires_at":   otpExpiresAt,
			"last_otp_sent_at": currentTime, // Cập nhật thời gian gửi OTP
			"otp_used":         false,
		}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set OTP"})
		return
	}

	err = utils.SendOTP(email, otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email", "otp_expires_in": 15 * 60}) // OTP có hiệu lực trong 15 phút
}

// Hàm xác minh OTP
func VerifyOTP(c *gin.Context) {
	email := c.PostForm("email")
	otp := c.PostForm("otp")

	collection := config.GetCollection("users")
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	currentTime := time.Now().Unix()
	if user.OTP != otp {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}

	if currentTime > user.OTPExpiresAt {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP has expired"})
		return
	}

	if user.OtpUsed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP has already been used"})
		return
	}

	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"email": email},
		bson.M{"$set": bson.M{"otp_used": true}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update OTP usage status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified, proceed to reset password"})
}

// Hàm đặt lại mật khẩu chỉ cho phép một lần
func ResetPassword(c *gin.Context) {
	email := c.PostForm("email")
	newPassword := c.PostForm("new_password")

	collection := config.GetCollection("users")
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	if user.PasswordResetCount >= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can only reset your password once"})
		return
	}

	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"email": email},
		bson.M{"$set": bson.M{
			"password":             newPassword,
			"otp":                  "",
			"otp_expires_at":       0,
			"otp_used":             true,
			"password_reset_count": user.PasswordResetCount + 1,
		}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
