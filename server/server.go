package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/Ashu23042000/grpc-practice/proto"
	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

// implementation for grpc service and to register this type(struct) with grpc we have to embeed Unimplemented indside this type(struct)
type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, in *pb.User) (*pb.CreatedUser, error) {
	newCreatedUser := &pb.CreatedUser{Name: in.GetName(), Email: in.GetEmail()}
	// s.user_list.Users = append(s.user_list.Users, newCreatedUser)
	return newCreatedUser, nil
}

func (s *Server) GreetUser(in *pb.User, stream pb.UserService_GreetUserServer) error {
	for i := 0; i < 5; i++ {
		res := fmt.Sprintf("Hello %s, for %d time", in.GetName(), i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}

func (s *Server) GreetMany(stream pb.UserService_GreetManyServer) error {
	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while getting stream from client: %v", err)
		}

		res += fmt.Sprintf("Hello %s \n", req.GetName())
	}
}

func (s *Server) GreetManyWithMultipleGreet(stream pb.UserService_GreetManyWithMultipleGreetServer) error {

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while getting stream from client: %v", err)
		}

		// res += fmt.Sprintf("Hello %s\n", req.GetName())
		res += "Hello " + req.GetName() + "\n"

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending stream to client: %v", err)
		}
	}

}

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
