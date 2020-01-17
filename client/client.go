package main

import (
	"encoding/json"
	"fmt"
	"github.com/Botinoc/task/connection"
	pb "github.com/Botinoc/task/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
)

const (
	address         = "localhost:50051"
	defaultFilename = "laziness.json"
)

//Функция парсит переданный фаил
func parseFile(file string) (*pb.Laziness, error) {
	var laziness *pb.Laziness
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &laziness)
	return laziness, err
}


func main() {
	// Установим соединение с сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не могу подключиться: %v", err)
	}
	defer conn.Close()
	client := pb.NewLazinessServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	laziness, err := parseFile(file)
	if err != nil {
		log.Fatalf("Не возможно распарсить фаил: %v", err)
	}

	r, err := client.CreateLaziness(context.Background(), laziness)
	if err != nil {
		log.Fatalf("Не удалось создать лень: %v", err)
	}

	log.Printf("Создана лень: %t", r.Created)

	//надо исправить дабы было привально
	//var index *pb.Index
	//resp, err := client.GetById(context.Background(), index)
	//if err != nil {
	//	log.Fatalf("Лень не найдена: %v", err)
	//}
	//fmt.Printf("Id: %v\n", resp.Laziness.GetId())
	//fmt.Printf("Description: %v\n", resp.Laziness.GetDescription())
	//fmt.Printf("Power of laziness: %d\n", resp.Laziness.GetPower())

	db := connection.GetConnection()
	rows, err := db.Query("select * from Laziness")
	for rows.Next(){
		fmt.Println(rows)
	}

}