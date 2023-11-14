package handler

import (
	"context"

	commiteepb "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"
	"google.golang.org/grpc/status"
)

// CreateSchedule retrieves a commitee request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
func (h *CommiteeHdl) CreateSchedule(ctx context.Context, req *commiteepb.CreateScheduleRequest) (*commiteepb.CreateScheduleResponse, error) {
	r, err := validateAndConvertSchedule(req.Schedule)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	s, err := h.Repository.CreateSchedule(ctx, r)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	room, err := h.Repository.GetRoom(ctx, s.RoomID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var tlRes []*commiteepb.TimeSlotsResponse
	tl, err := h.Repository.GetTimeSlotsByScheduleID(ctx, s.ID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	for _, t := range tl {
		tlRes = append(tlRes, &commiteepb.TimeSlotsResponse{
			Id:         int64(t.ID),
			ScheduleID: int64(t.ScheduleID),
		})
	}

	resp := &commiteepb.CreateScheduleResponse{
		Schedule: &commiteepb.ScheduleResponse{
			TimeSlots: tlRes,
			Room: &commiteepb.RoomSchedule{
				Id:          int64(room.ID),
				Name:        room.Name,
				Type:        room.Type,
				School:      room.School,
				Description: room.Description,
			},
			ThesisID: int64(s.RoomID),
		},
	}

	return resp, nil
}

// GetSchedule returns a commitee in db given by id
func (h *CommiteeHdl) GetSchedule(ctx context.Context, req *commiteepb.GetScheduleRequest) (*commiteepb.GetScheduleResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	s, err := h.Repository.GetSchedule(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	room, err := h.Repository.GetRoom(ctx, s.RoomID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var tlRes []*commiteepb.TimeSlotsResponse
	tl, err := h.Repository.GetTimeSlotsByScheduleID(ctx, s.ID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	for _, t := range tl {
		tlRes = append(tlRes, &commiteepb.TimeSlotsResponse{
			Id:         int64(t.ID),
			ScheduleID: int64(t.ScheduleID),
		})
	}

	resp := &commiteepb.GetScheduleResponse{
		Schedule: &commiteepb.ScheduleResponse{
			TimeSlots: tlRes,
			Room: &commiteepb.RoomSchedule{
				Id:          int64(room.ID),
				Name:        room.Name,
				Type:        room.Type,
				School:      room.School,
				Description: room.Description,
			},
			ThesisID: int64(s.RoomID),
		},
	}

	return resp, nil
}

func (h *CommiteeHdl) GetSchedules(ctx context.Context, req *commiteepb.GetSchedulesRequest) (*commiteepb.GetSchedulesResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	rs, _, err := h.Repository.GetSchedules(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var sRes []*commiteepb.ScheduleResponse
	for _, s := range rs {
		room, err := h.Repository.GetRoom(ctx, s.RoomID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		var tlRes []*commiteepb.TimeSlotsResponse
		tl, err := h.Repository.GetTimeSlotsByScheduleID(ctx, s.ID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		for _, t := range tl {
			tlRes = append(tlRes, &commiteepb.TimeSlotsResponse{
				Id:         int64(t.ID),
				ScheduleID: int64(t.ScheduleID),
			})
		}

		sRes = append(sRes, &commiteepb.ScheduleResponse{
			TimeSlots: tlRes,
			Room: &commiteepb.RoomSchedule{
				Id:          int64(room.ID),
				Name:        room.Name,
				Type:        room.Type,
				School:      room.School,
				Description: room.Description,
			},
			ThesisID: int64(s.RoomID),
		},
		)
	}

	return &commiteepb.GetSchedulesResponse{
		Schedules: sRes,
	}, nil
}

func (h *CommiteeHdl) GetSchedulesByThesisID(ctx context.Context, req *commiteepb.GetSchedulesByThesisIDRequest) (*commiteepb.GetSchedulesResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	rs, err := h.Repository.GetSchedulesByThesisID(ctx, int(req.ThesisID))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var rsRes []*commiteepb.ScheduleResponse
	for _, s := range rs {
		room, err := h.Repository.GetRoom(ctx, s.RoomID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		var tlRes []*commiteepb.TimeSlotsResponse
		tl, err := h.Repository.GetTimeSlotsByScheduleID(ctx, s.ID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		for _, t := range tl {
			tlRes = append(tlRes, &commiteepb.TimeSlotsResponse{
				Id:         int64(t.ID),
				ScheduleID: int64(t.ScheduleID),
			})
		}

		rsRes = append(rsRes, &commiteepb.ScheduleResponse{
			TimeSlots: tlRes,
			Room: &commiteepb.RoomSchedule{
				Id:          int64(room.ID),
				Name:        room.Name,
				Type:        room.Type,
				School:      room.School,
				Description: room.Description,
			},
			ThesisID: int64(s.RoomID),
		},
		)
	}

	return &commiteepb.GetSchedulesResponse{
		Schedules: rsRes,
	}, nil
}

func validateAndConvertSchedule(pbSchedule *commiteepb.ScheduleInput) (repository.ScheduleInputRepo, error) {
	if err := pbSchedule.Validate(); err != nil {
		return repository.ScheduleInputRepo{}, err
	}

	return repository.ScheduleInputRepo{
		RoomID:   int(pbSchedule.RoomID),
		ThesisID: int(pbSchedule.ThesisID),
	}, nil
}
