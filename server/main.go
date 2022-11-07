package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	pb "github.com/irhamsahbana/simple-grpc/student"
	"google.golang.org/grpc"
)

type DataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu      sync.Mutex
	student []*pb.Student
}

func (s *DataStudentServer) FindStudentByEmail(ctx context.Context, in *pb.Student) (*pb.Student, error) {
	fmt.Println("Incoming request find student by email from client ...")
	for _, student := range s.student {
		if student.Email == in.Email {
			return student, nil
		}
	}
	return nil, fmt.Errorf("student with email %s not found", in.Email)
}

func (s *DataStudentServer) LoadData() {
	data, err := ioutil.ReadFile("data/students.json")
	if err != nil {
		log.Fatalln("error in read file", err)
	}

	if err := json.Unmarshal(data, &s.student); err != nil {
		log.Fatalln("error in unmarshal", err)
	}
}

func newServer() *DataStudentServer {
	s := &DataStudentServer{}
	s.LoadData()
	return s
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("error in listen", err)
	}

	s := newServer()
	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error in serve", err)
	}
}
