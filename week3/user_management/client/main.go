package main

import (
	pb "GrabAssignments/week3/user_management"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5501"
)

func test(client pb.FeedbackManagementClient) {
	fb1 := pb.PassengerFeedback{
		PassengerID: 1,
		BookingCode: "A",
		Feedback:    "feedback 1"}
	fb11 := pb.PassengerFeedback{
		PassengerID: 2,
		BookingCode: "A",
		Feedback:    "feedback 11"}
	fb2 := pb.PassengerFeedback{
		PassengerID: 1,
		BookingCode: "B",
		Feedback:    "feedback 2"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Successful
	_, err := client.AddFeedBack(ctx, &pb.AddFeedbackRequest{PassengerFeedback: &fb1})
	if err == nil {
		log.Println("Added fb1")
	} else {
		log.Println("Could not add fb1, reason: ", err)
	}

	//Unsuccessful due to duplicated booking code
	_, err = client.AddFeedBack(ctx, &pb.AddFeedbackRequest{PassengerFeedback: &fb11})
	if err == nil {
		log.Println("Added fb11")
	} else {
		log.Println("Could not add fb11, reason: ", err)
	}

	//Successful
	_, err = client.AddFeedBack(ctx, &pb.AddFeedbackRequest{PassengerFeedback: &fb2})
	if err == nil {
		log.Println("Added fb2")
	} else {
		log.Println("Could not add fb2, reason: ", err)
	}

	//Return 1 object
	fb, err := client.GetByBookingCode(ctx, &pb.GetByBookingCodeRequest{BookingCode: fb1.BookingCode})
	log.Printf("Passenger feedback with booking code %s: %s", fb1.BookingCode, fb)

	//Return 1 object
	fb, err = client.GetByBookingCode(ctx, &pb.GetByBookingCodeRequest{BookingCode: fb2.BookingCode})
	log.Printf("Passenger feedback with booking code %s: %s", fb2.BookingCode, fb)

	//Return 2 objects
	fbList, err := client.GetByPassengerId(ctx, &pb.GetByPassengerIdRequest{PassengerID: fb1.PassengerID})
	log.Printf("Passenger feedbacks with passenger id %d: %s", fb1.PassengerID, fbList)

	_, err = client.DeleteByPassengerId(ctx, &pb.DeleteByPassengerIdRequest{PassengerID: fb1.PassengerID})
	if err == nil {
		log.Printf("Deleted feedback(s) with passenger id %d", fb1.PassengerID)
	} else {
		log.Println(err)
	}

	//Unsucessful because of empty list
	fbList, err = client.GetByPassengerId(ctx, &pb.GetByPassengerIdRequest{PassengerID: fb1.PassengerID})
	log.Printf("Passenger feedback with passenger id %d: %s", fb1.PassengerID, fbList)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot connect to server")
	}
	defer conn.Close()

	client := pb.NewFeedbackManagementClient(conn)
	test(client)
}
