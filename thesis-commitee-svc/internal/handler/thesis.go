package handler

import (
	"context"
	"log"

	thesispb "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	"google.golang.org/grpc/status"
)

// // CreateThesis retrieves a thesis request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
// func (h *CommiteeHdl) CreateThesis(ctx context.Context, req *thesispb.CreateThesisRequest) (*thesispb.CreateThesisResponse, error) {
// 	log.Println("calling insert thesis...")
// 	p, err := validateAndConvertThesis(req.Thesis)
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}
// 	cOut, err := h.Repository.CreateThesis(ctx, p)
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	resp := &thesispb.CreateThesisResponse{
// 		Response: &thesispb.CommonThesisResponse{
// 			StatusCode: 201,
// 			Message:    "Created",
// 		},
// 		Thesis: &thesispb.ThesisResponse{
// 			Id: int64(cOut.ID),
// 		},
// 	}

// 	return resp, nil
// }
//
// // GetThesis returns a thesis in db given by id
// func (h *CommiteeHdl) GetThesis(ctx context.Context, req *thesispb.GetThesisRequest) (*thesispb.GetThesisResponse, error) {
// 	log.Println("calling get thesis...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}
// 	p, err := h.Repository.GetThesis(ctx, int(req.GetId()))
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	pResp := thesispb.ThesisResponse{
// 		Id: int64(p.ID),
// 	}

// 	resp := &thesispb.GetThesisResponse{
// 		Response: &thesispb.CommonThesisResponse{
// 			StatusCode: 200,
// 			Message:    "OK",
// 		},
// 		Thesis: &pResp,
// 	}
// 	return resp, nil
// }

func (h *CommiteeHdl) GetThesises(ctx context.Context, req *thesispb.GetThesisesRequest) (*thesispb.GetThesisesResponse, error) {
	log.Println("calling get all thesiss...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, _, err := h.Repository.GetThesiss(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*thesispb.ThesisResponse
	for _, p := range ps {
		schedule, err := h.Repository.GetSchedule(ctx, p.ID)
		if err != nil {
			return nil, err
		}

		timeSlots, err := h.Repository.GetTimeSlotsByScheduleID(ctx, schedule.ID)
		if err != nil {
			return nil, err
		}

		var timeSlotRes []*thesispb.TimeSlotsResponse
		for _, tl := range timeSlots {
			timeSlot, err := h.Repository.GetCommiteeByTimeSlotsID(ctx, tl.ID)
			if err != nil {
				return nil, err
			}

			timeSlotRes = append(timeSlotRes, &thesispb.TimeSlotsResponse{
				TimeSlot: &thesispb.TimeSlot{
					Date:  timeSlot.StartDate.String(),
					Shift: timeSlot.Period,
					Id:    int64(timeSlot.ID),
					Time:  timeSlot.Time,
				},
				Id:         int64(tl.ID),
				ScheduleID: int64(tl.ScheduleID),
			})
		}

		room, err := h.Repository.GetRoom(ctx, schedule.RoomID)
		if err != nil {
			return nil, err
		}

		councils, _, err := h.Repository.GetCouncilsByThesisID(ctx, p.ID)
		if err != nil {
			return nil, err
		}

		var councilRes []string
		for _, c := range councils {
			councilRes = append(councilRes, c.LecturerID)
		}

		psResp = append(psResp, &thesispb.ThesisResponse{
			Thesis: &thesispb.Thesis{
				Schedule: &thesispb.ScheduleResponse{
					TimeSlots: timeSlotRes,
					Room: &thesispb.RoomSchedule{
						Id:          int64(room.ID),
						Name:        room.Name,
						Type:        room.Type,
						School:      room.School,
						Description: room.Description,
					},
				},
				CouncilId: councilRes,
				Id:        int64(p.ID),
			},
		})
	}

	return &thesispb.GetThesisesResponse{
		ScheduleReport: psResp[0],
	}, nil
}

// func validateAndConvertThesis(pbThesis *thesispb.Thesis) (repository.ThesisInputRepo, error) {
// 	if err := pbThesis.Validate(); err != nil {
// 		return repository.ThesisInputRepo{}, err
// 	}

// 	return repository.ThesisInputRepo{

// 	}, nil
// }
