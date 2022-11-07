package main

import (
	"context"
	"log"
	"time"

	pb "github.com/irhamsahbana/simple-grpc/student"
	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := &pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, s)
	if err != nil {
		log.Fatalf("error in find student by email => %v", err)
	}

	log.Printf("student => %v", student)
}

func main() {
	var opts []grpc.DialOption

	// Set up a connection to the server.
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":50051", opts...)
	if err != nil {
		log.Fatalf("did not connect, error in dial: %v", err)
	}

	defer conn.Close()
	client := pb.NewDataStudentClient(conn)

	getDataStudentByEmail(client, "john@example.com")
	getDataStudentByEmail(client, "jane@example.com")
	getDataStudentByEmail(client, "")
}
