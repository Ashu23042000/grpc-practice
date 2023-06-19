package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/Ashu23042000/grpc-practice/proto"
)

const (
	address = "localhost:5000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Unary request -------------------------------------------------

	fmt.Println("\nUnary request ----------------------------------------------\n")

	res, err := c.CreateUser(ctx, &pb.User{Name: "Sagar", Email: "ashu23042000@gmail.com"})

	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	log.Printf("Message from server: %v", res)

	// server streaming------------------------------------------

	fmt.Println("\nserver streaming -------------------------------------\n")

	param := &pb.User{
		Name:  "Ashu",
		Email: "ashu23042000@gmail.com",
	}

	stream, err := c.GreetUser(context.Background(), param)

	if err != nil {
		log.Fatalf("Failed to greet user: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while getting stream from server: %v", err)
		}

		log.Println(msg.Result)
	}

	// client streaming--------------------------------------

	fmt.Println("\nclient streaming ----------------------------------\n")

	params := []*pb.User{
		{Name: "Ashu", Email: "ashu@gmail.com"},
		{Name: "Sagar", Email: "sagargmail.com"},
		{Name: "Saggy Bhai", Email: "saggy@gmail.com"},
	}

	streamm, err := c.GreetMany(context.Background())

	if err != nil {
		log.Fatalf("Failed to greet many users: %v", err)
	}

	for _, req := range params {
		streamm.Send(req)
	}

	ress, err := streamm.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while getting response from server: %v", err)
	}

	log.Println(ress)

	// bi-directional streaming---------------------------

	fmt.Println("\nbi-directional streaming --------------------------------------\n")

	streammm, err := c.GreetManyWithMultipleGreet(context.Background())

	if err != nil {
		log.Fatalf("Failed to greet : %v", err)
	}

	chann := make(chan struct{})

	// goroutine for sending stream to server-----------

	go func() {
		for _, req := range params {
			streammm.Send(req)
		}
		streammm.CloseSend()
	}()

	// goroutine for recieving stream from server----------------

	go func() {
		for {
			res, err := streammm.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while getting stream from server: %v", err)
			}

			log.Println(res.Result)
		}

		close(chann)
	}()
	<-chann
}
