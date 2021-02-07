package main

import (
	"os"

	"github.com/iamtheluiz/codepix/codepix-go/application/grpc"
	"github.com/iamtheluiz/codepix/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
