package main

import (
	"context"
	"log"
	"net"

	"github.com/yourusername/todo-app/user-service/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	proto.UnimplementedUserServiceServer
	db *gorm.DB
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func (s *server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	user := User{Username: req.GetUsername(), Password: req.GetPassword()}
	result := s.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &proto.RegisterResponse{Id: uint64(user.ID)}, nil
}

func (s *server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var user User
	result := s.db.Where("username = ? AND password = ?", req.GetUsername(), req.GetPassword()).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// Генерация JWT токена (упрощённый пример)
	token := "some-jwt-token"
	return &proto.LoginResponse{Token: token}, nil
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=user_service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &server{db: db})
	log.Println("User Service is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
