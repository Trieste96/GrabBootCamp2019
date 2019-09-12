package main

import (
	pb "GrabBootCamp2019/week3/user_management"
	"errors"
)

//FeedbackStorage stores feedbacks by booking code and by passenger id
type FeedbackStorage struct {
	BookingCodeMap       map[string]*pb.PassengerFeedback
	PassengerFeedbackMap map[int32][]*pb.PassengerFeedback
}

func (s *FeedbackStorage) addNew(req *pb.AddFeedbackRequest) error {
	fb := req.PassengerFeedback
	if _, ok := s.BookingCodeMap[fb.BookingCode]; ok {
		return errors.New("Duplicate booking code")
	}
	s.BookingCodeMap[fb.BookingCode] = fb
	s.PassengerFeedbackMap[fb.PassengerID] = append(s.PassengerFeedbackMap[fb.PassengerID], fb)
	return nil
}

func (s *FeedbackStorage) getByBookingCode(req *pb.GetByBookingCodeRequest) (*pb.PassengerFeedback, error) {
	if fb, ok := s.BookingCodeMap[req.BookingCode]; ok {
		return fb, nil
	}
	return nil, errors.New("Passenger feedback not found")
}

func (s *FeedbackStorage) getByPassengerID(req *pb.GetByPassengerIdRequest) ([]*pb.PassengerFeedback, error) {
	if fbList, ok := s.PassengerFeedbackMap[req.PassengerID]; ok {
		return fbList, nil
	}
	return nil, errors.New("Passenger feedback not found")
}

func (s *FeedbackStorage) deleteByPassengerID(req *pb.DeleteByPassengerIdRequest) error {
	if fbList, ok := s.PassengerFeedbackMap[req.PassengerID]; ok {
		for _, v := range fbList {
			delete(s.BookingCodeMap, v.BookingCode)
		}
		delete(s.PassengerFeedbackMap, req.PassengerID)
		return nil
	}
	return errors.New("Passenger feedback not found")
}
