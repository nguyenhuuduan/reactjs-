package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Email              string             `bson:"email"`
	Password           string             `bson:"password"`
	OTP                string             `bson:"otp,omitempty"`
	OTPExpiresAt       int64              `bson:"otp_expires_at,omitempty"`
	OtpUsed            bool               `bson:"otp_used"`             // Đánh dấu OTP đã được dùng hay chưa
	PasswordResetCount int                `bson:"password_reset_count"` // Theo dõi số lần đặt lại mật khẩu
	LastOTPSentAt      int64              `bson:"last_otp_sent_at"`     // Thêm trường này
	Profile            Profile            `bson:"profile"`              // Thông tin cá nhân và hẹn hò
}

type Profile struct {
	Name     string `bson:"name"`
	Birthday string `bson:"birthday"`
	Gender   string `bson:"gender"`
	Location string `bson:"location"`
	Bio      string `bson:"bio"`
	// Bạn có thể thêm các trường khác liên quan đến hẹn hò ở đây (tương tự như Bumble)
}
