package handler

import (
	"context"
	"log"

	commiteepb "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"
	"google.golang.org/grpc/status"
)

// CreateRoom retrieves a commitee request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
func (h *CommiteeHdl) CreateRoom(ctx context.Context, req *commiteepb.CreateRoomRequest) (*commiteepb.CreateRoomResponse, error) {
	r, err := validateAndConvertRoom(req.Room)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	rOut, err := h.Repository.CreateRoom(ctx, r)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &commiteepb.CreateRoomResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 201,
			Message:    "Created",
		},
		Room: &commiteepb.RoomResponse{
			Id:          int64(rOut.ID),
			Name:        rOut.Name,
			Type:        rOut.Type,
			School:      rOut.School,
			Description: rOut.Description,
		},
	}

	return resp, nil
}

// GetRoom returns a commitee in db given by id
func (h *CommiteeHdl) GetRoom(ctx context.Context, req *commiteepb.GetRoomRequest) (*commiteepb.GetRoomResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	r, err := h.Repository.GetRoom(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := commiteepb.RoomResponse{
		Id:          int64(r.ID),
		Name:        r.Name,
		Type:        r.Type,
		School:      r.School,
		Description: r.Description,
	}

	resp := &commiteepb.GetRoomResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Room: &pResp,
	}
	return resp, nil
}

func (c *CommiteeHdl) UpdateRoom(ctx context.Context, req *commiteepb.UpdateRoomRequest) (*commiteepb.UpdateRoomResponse, error) {
	log.Println("calling update commitee...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	r, err := validateAndConvertRoom(req.Room)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Repository.UpdateRoom(ctx, int(req.GetId()), repository.RoomInputRepo{
		Name:        r.Name,
		Type:        r.Type,
		School:      r.School,
		Description: r.Description,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &commiteepb.UpdateRoomResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *CommiteeHdl) DeleteRoom(ctx context.Context, req *commiteepb.DeleteRoomRequest) (*commiteepb.DeleteRoomResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Repository.DeleteRoom(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &commiteepb.DeleteRoomResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *CommiteeHdl) GetRooms(ctx context.Context, req *commiteepb.GetRoomsRequest) (*commiteepb.GetRoomsResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	nameParam := req.GetName()
	typeParam := req.GetType()
	schoolParam := req.GetSchool()
	rs, count, err := h.Repository.GetRooms(ctx, repository.RoomFilter{
		Name:   &nameParam,
		Type:   &typeParam,
		School: &schoolParam,
	})
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var rsResp []*commiteepb.RoomResponse
	for _, r := range rs {
		rsResp = append(rsResp, &commiteepb.RoomResponse{
			Id:          int64(r.ID),
			Name:        r.Name,
			Type:        r.Type,
			School:      r.School,
			Description: r.Description,
		})
	}

	return &commiteepb.GetRoomsResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Rooms:      rsResp,
		TotalCount: int64(count),
	}, nil
}

func validateAndConvertRoom(pbRoom *commiteepb.RoomInput) (repository.RoomInputRepo, error) {
	if err := pbRoom.Validate(); err != nil {
		return repository.RoomInputRepo{}, err
	}

	return repository.RoomInputRepo{
		Name:        pbRoom.Name,
		Type:        pbRoom.Type,
		School:      pbRoom.School,
		Description: pbRoom.Description,
	}, nil
}
