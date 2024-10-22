package controllers

import (
	"context"
	"dating_app/config"
	"dating_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Lấy thông tin người dùng dựa vào ID
func GetUserProfile(c *gin.Context) {
	id := c.Param("id") // Lấy ID từ URL

	// Chuyển đổi id thành ObjectID của MongoDB
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Tìm người dùng theo ID
	var user models.User
	collection := config.GetCollection("users")
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Trả về thông tin cá nhân của người dùng
	c.JSON(http.StatusOK, gin.H{
		"email":   user.Email,
		"profile": user.Profile,
	})
}

// Cập nhật thông tin người dùng dựa vào ID
func UpdateUserProfile(c *gin.Context) {
	id := c.Param("id") // Lấy ID từ URL

	// Chuyển đổi id thành ObjectID của MongoDB
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Lấy dữ liệu thông tin cập nhật từ body
	var updatedProfile models.Profile
	if err := c.ShouldBindJSON(&updatedProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	collection := config.GetCollection("users")

	// Tìm và cập nhật thông tin người dùng theo ID
	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"profile": updatedProfile}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
