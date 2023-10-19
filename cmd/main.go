package main

import (
	"log"
	"os"

	"github.com/codeedu/imersao/codepix-go/application/grpc"
	"github.com/codeedu/imersao/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	env := os.Getenv("env")

	log.Printf("Environment: %s", env)

	database = db.ConnectDB(env)
	grpc.StartGrpcServer(database, 50051)
}
