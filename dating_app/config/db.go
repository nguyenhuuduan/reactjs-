package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDatabase() {
	// Tạo tùy chọn kết nối cho MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Kết nối tới MongoDB bằng cách dùng trực tiếp mongo.Connect
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error creating MongoDB client:", err)
		return
	}

	// Đặt timeout cho quá trình kết nối và Ping để kiểm tra kết nối
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Kiểm tra xem kết nối có thành công không bằng cách "ping" MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	DB = client
	fmt.Println("Connected to MongoDB!")
}

// Lấy một collection từ cơ sở dữ liệu MongoDB
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("dating_app").Collection(collectionName)
}
