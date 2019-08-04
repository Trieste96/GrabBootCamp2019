package main

import (
	"context"
	"log"
	"net"

	pb "GrabAssignments/week3/user_management"

	"google.golang.org/grpc"
)

const (
	port = ":5501"
)

//Server implements FeedbackManagementServer interface
type Server struct{}

func (s *Server) AddFeedBack(ctx context.Context, req *pb.AddFeedbackRequest) (*pb.FeedbackResponse, error) {
	log.Println("Receveid %s", req)
	return &pb.FeedbackResponse{Success: true}, nil
}

func (s *Server) GetByPassengerId(ctx context.Context, req *pb.GetByPassengerIdRequest) (*pb.FeedbackListResponse, error) {
	log.Println("Receveid %s", req)
	return &pb.FeedbackListResponse{FeedbackList: nil}, nil
}

func (s *Server) GetByBookingCode(ctx context.Context, req *pb.GetByBookingCodeRequest) (*pb.FeedbackListResponse, error) {
	log.Println("Receveid %s", req)
	return &pb.FeedbackListResponse{FeedbackList: nil}, nil
}

func (s *Server) DeleteByPassengerId(ctx context.Context, req *pb.DeleteByPassengerIdRequest) (*pb.FeedbackResponse, error) {
	log.Println("Receveid %s", req)
	return &pb.FeedbackResponse{Success: true}, nil
}

func main() {
	log.Println("Start main")
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Port is occupied")
	}

	s := grpc.NewServer()
	pb.RegisterFeedbackManagementServer(s, &Server{})
	if err := s.Serve(conn); err != nil {
		log.Fatal("Failed to create server")
	}
}
