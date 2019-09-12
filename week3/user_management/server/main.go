package main

import (
	"context"
	"log"
	"net"

	pb "GrabBootCamp2019/week3/user_management"

	"google.golang.org/grpc"
)

const (
	port = ":5501"
)

var store = FeedbackStorage{
	BookingCodeMap:       make(map[string]*pb.PassengerFeedback),
	PassengerFeedbackMap: make(map[int32][]*pb.PassengerFeedback)}

//Server implements FeedbackManagementServer interface
type Server struct{}

func (s *Server) AddFeedBack(ctx context.Context, req *pb.AddFeedbackRequest) (*pb.FeedbackResponse, error) {
	log.Printf("AddFeedBack %v\n", req.PassengerFeedback)
	err := store.addNew(req)
	if err != nil {
		return nil, err
	}
	return &pb.FeedbackResponse{Success: true}, nil
}

func (s *Server) GetByPassengerId(ctx context.Context, req *pb.GetByPassengerIdRequest) (*pb.FeedbackListResponse, error) {
	log.Printf("GetByPassengerId %v\n", req.PassengerID)
	fbList, err := store.getByPassengerID(req)
	if err != nil {
		return nil, err
	}
	return &pb.FeedbackListResponse{FeedbackList: fbList, Success: true}, nil
}

func (s *Server) GetByBookingCode(ctx context.Context, req *pb.GetByBookingCodeRequest) (*pb.FeedbackResponse, error) {
	log.Printf("GetByBookingCode %v\n", req.BookingCode)
	fb, err := store.getByBookingCode(req)
	if err != nil {
		return nil, err
	}
	return &pb.FeedbackResponse{PassengerFeedback: fb, Success: true}, nil
}

func (s *Server) DeleteByPassengerId(ctx context.Context, req *pb.DeleteByPassengerIdRequest) (*pb.FeedbackResponse, error) {
	log.Printf("DeleteByPassengerId %v\n", req.PassengerID)
	err := store.deleteByPassengerID(req)
	if err != nil {
		return nil, err
	}
	return &pb.FeedbackResponse{Success: true}, nil
}

func main() {
	log.Println("Start main")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Port is occupied")
	}

	s := grpc.NewServer()
	pb.RegisterFeedbackManagementServer(s, &Server{})
	if err := s.Serve(listen); err != nil {
		log.Fatal("Failed to create server")
	}
	defer listen.Close()
	pb.RegisterFeedbackManagementServer(s)
}
