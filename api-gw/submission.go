package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	submissionSvcV1 "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
)

type submissionServiceGW struct {
	pb.UnimplementedSubmissionServiceServer
	submissionClient submissionSvcV1.SubmissionServiceClient
	classroomClient  classroomSvcV1.ClassroomServiceClient
	exerciseClient   exerciseSvcV1.ExerciseServiceClient
}

func NewSubmissionsService(submissionClient submissionSvcV1.SubmissionServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, exerciseClient exerciseSvcV1.ExerciseServiceClient) *submissionServiceGW {
	return &submissionServiceGW{
		submissionClient: submissionClient,
		classroomClient:  classroomClient,
		exerciseClient:   exerciseClient,
	}
}

func (u *submissionServiceGW) CreateSubmission(ctx context.Context, req *pb.CreateSubmissionRequest) (*pb.CreateSubmissionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	exists, err := u.exerciseClient.GetExercise(ctx, &exerciseSvcV1.GetExerciseRequest{Id: req.GetSubmission().GetExerciseID()})
	if err != nil {
		return nil, err
	}

	if exists.GetResponse().StatusCode == 404 {
		return &pb.CreateSubmissionResponse{
			Response: &pb.CommonSubmissionResponse{
				StatusCode: 404,
				Message:    "Exercise does not exist",
			},
		}, nil
	}

	res, err := u.submissionClient.CreateSubmission(ctx, &submissionSvcV1.CreateSubmissionRequest{
		Submission: &submissionSvcV1.SubmissionInput{
			UserID:         req.GetSubmission().GetUserID(),
			ExerciseID:     req.GetSubmission().GetExerciseID(),
			SubmissionDate: req.GetSubmission().GetSubmissionDate(),
			Status:         req.GetSubmission().GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *submissionServiceGW) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionResponse, error) {
	res, err := u.submissionClient.GetSubmission(ctx, &submissionSvcV1.GetSubmissionRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Submission: &pb.SubmissionResponse{
			Id:             res.GetSubmission().GetId(),
			UserID:         res.GetSubmission().GetUserID(),
			ExerciseID:     res.GetSubmission().GetExerciseID(),
			SubmissionDate: res.GetSubmission().GetSubmissionDate(),
			Status:         res.GetSubmission().GetStatus(),
		},
	}, nil
}

func (u *submissionServiceGW) UpdateSubmission(ctx context.Context, req *pb.UpdateSubmissionRequest) (*pb.UpdateSubmissionResponse, error) {
	exists, err := u.exerciseClient.GetExercise(ctx, &exerciseSvcV1.GetExerciseRequest{Id: req.GetSubmission().GetExerciseID()})
	if err != nil {
		return nil, err
	}

	if exists.GetResponse().StatusCode == 404 {
		return &pb.UpdateSubmissionResponse{
			Response: &pb.CommonSubmissionResponse{
				StatusCode: 404,
				Message:    "Exercise does not exist",
			},
		}, nil
	}

	res, err := u.submissionClient.UpdateSubmission(ctx, &submissionSvcV1.UpdateSubmissionRequest{
		Id: req.GetId(),
		Submission: &submissionSvcV1.SubmissionInput{
			UserID:         req.GetSubmission().GetUserID(),
			ExerciseID:     req.GetSubmission().GetExerciseID(),
			SubmissionDate: req.GetSubmission().GetSubmissionDate(),
			Status:         req.GetSubmission().GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *submissionServiceGW) DeleteSubmission(ctx context.Context, req *pb.DeleteSubmissionRequest) (*pb.DeleteSubmissionResponse, error) {
	res, err := u.submissionClient.DeleteSubmission(ctx, &submissionSvcV1.DeleteSubmissionRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *submissionServiceGW) GetAllSubmissionsOfExercise(ctx context.Context, req *pb.GetAllSubmissionsOfExerciseRequest) (*pb.GetAllSubmissionsOfExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.submissionClient.GetAllSubmissionsOfExercise(ctx, &submissionSvcV1.GetAllSubmissionsOfExerciseRequest{})
	if err != nil {
		return nil, err
	}

	var submissions []*pb.SubmissionResponse
	for _, p := range res.GetSubmissions() {
		submissions = append(submissions, &pb.SubmissionResponse{
			Id:             p.Id,
			UserID:         p.UserID,
			ExerciseID:     p.ExerciseID,
			SubmissionDate: p.SubmissionDate,
			Status:         p.Status,
		})
	}

	return &pb.GetAllSubmissionsOfExerciseResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount:  res.GetTotalCount(),
		Submissions: submissions,
	}, nil
}
