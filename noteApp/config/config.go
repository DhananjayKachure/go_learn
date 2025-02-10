package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	clientoptons := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientoptons)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Database Connection Failed!")
	} else {
		fmt.Println("Database Connected Successfully!")
	}

	DB = client.Database(os.Getenv("DB_NAME"))

}
func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
