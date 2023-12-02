package service

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	roomSvcV1 "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
)

type roomServiceGW struct {
	pb.UnimplementedRoomServiceServer
	roomClient roomSvcV1.CommiteeServiceClient
}

func NewRoomsService(roomClient roomSvcV1.CommiteeServiceClient) *roomServiceGW {
	return &roomServiceGW{
		roomClient: roomClient,
	}
}

func (u *roomServiceGW) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.roomClient.CreateRoom(ctx, &roomSvcV1.CreateRoomRequest{
		Room: &roomSvcV1.RoomInput{
			Name:        req.GetRoom().Name,
			Type:        req.GetRoom().Type,
			School:      req.GetRoom().School,
			Description: req.GetRoom().Description,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoomResponse{
		Response: &pb.CommonRoomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *roomServiceGW) GetRoom(ctx context.Context, req *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.roomClient.GetRoom(ctx, &roomSvcV1.GetRoomRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetRoomResponse{
		Response: &pb.CommonRoomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Room: &pb.RoomResponse{
			Id:          res.GetRoom().Id,
			Name:        res.GetRoom().Name,
			Type:        res.GetRoom().Type,
			School:      res.GetRoom().School,
			Description: res.GetRoom().Description,
		},
	}, nil
}

func (u *roomServiceGW) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*pb.UpdateRoomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.roomClient.UpdateRoom(ctx, &roomSvcV1.UpdateRoomRequest{
		Id: req.GetId(),
		Room: &roomSvcV1.RoomInput{
			Name:        req.GetRoom().Name,
			Type:        req.GetRoom().Type,
			School:      req.GetRoom().School,
			Description: req.GetRoom().Description,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateRoomResponse{
		Response: &pb.CommonRoomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *roomServiceGW) DeleteRoom(ctx context.Context, req *pb.DeleteRoomRequest) (*pb.DeleteRoomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.roomClient.DeleteRoom(ctx, &roomSvcV1.DeleteRoomRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteRoomResponse{
		Response: &pb.CommonRoomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *roomServiceGW) GetRooms(ctx context.Context, req *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	nameParam := req.GetName()
	typeParam := req.GetType()
	schoolParam := req.GetSchool()
	res, err := u.roomClient.GetRooms(ctx, &roomSvcV1.GetRoomsRequest{
		Name:   &nameParam,
		Type:   &typeParam,
		School: &schoolParam,
	})
	if err != nil {
		return nil, err
	}

	var rooms []*pb.RoomResponse
	for _, p := range res.GetRooms() {
		rooms = append(rooms, &pb.RoomResponse{
			Id:          p.Id,
			Name:        p.Name,
			Type:        p.Type,
			School:      p.School,
			Description: p.Description,
		})
	}

	return &pb.GetRoomsResponse{
		Response: &pb.CommonRoomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Rooms:      rooms,
	}, nil
}
