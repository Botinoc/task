package main

import (
	"fmt"
	"github.com/Botinoc/task/connection"
	"log"
	"net"
	pb "github.com/Botinoc/task/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//GetById временно не работает
func (s *service) GetById(ctx context.Context, index *pb.Index) (*pb.Response, error) {
	db := connection.GetConnection()
	rows, err := db.Query("select * from Laziness where id = 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	var laziness *pb.Laziness
	err = rows.Scan(laziness.Id, laziness.Description, laziness.Power)
	if err != nil {
		fmt.Println(err)
	}
	return &pb.Response{Laziness: laziness}, nil
}


type service struct {}


func (s *service) CreateLaziness(ctx context.Context, req *pb.Laziness) (*pb.Response, error) {
	db := connection.GetConnection()
	_, err := db.Exec("insert into Laziness (description, powerLaziness) values ($1, $2)", req.Description, req.Power)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Laziness: req}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterLazinessServiceServer(s, &service{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}