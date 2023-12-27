package service

import (
	"context"
	"errors"

	repository "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository/user"
)

type StudentDefInputSvc struct {
	UserID       string
	InstructorID string
	TimeSlotsID  int
}

type StudentDefOutputSvc struct {
	ID           int
	UserID       string
	InstructorID string
	TimeSlotsID  int
}

// CreateStudentDef creates a new member in db given by member model
func (s *UserSvc) CreateStudentDef(ctx context.Context, sd StudentDefInputSvc) error {
	sdRepo := repository.StudentDefInputRepo{
		UserID:       sd.UserID,
		InstructorID: sd.InstructorID,
		TimeSlotsID:  sd.TimeSlotsID,
	}

	if err := s.Repository.CreateStudentDef(ctx, sdRepo); err != nil {
		if errors.Is(err, repository.ErrStudentDefExisted) {
			return ErrStudentDefExisted
		}
		return err
	}

	return nil
}

// GetStudentDef returns a member in db given by id
func (s *UserSvc) GetStudentDef(ctx context.Context, id int) (StudentDefOutputSvc, error) {
	sd, err := s.Repository.GetStudentDef(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrStudentDefNotFound) {
			return StudentDefOutputSvc{}, ErrStudentDefNotFound
		}
		return StudentDefOutputSvc{}, err
	}

	return StudentDefOutputSvc{
		ID:           sd.ID,
		UserID:       sd.UserID,
		InstructorID: sd.InstructorID,
		TimeSlotsID:  sd.TimeSlotsID,
	}, nil
}

// UpdateStudentDef updates the specified member by id
func (s *UserSvc) UpdateStudentDef(ctx context.Context, id int, sd StudentDefInputSvc) error {
	if err := s.Repository.UpdateStudentDef(ctx, id, repository.StudentDefInputRepo{
		UserID:       sd.UserID,
		InstructorID: sd.InstructorID,
		TimeSlotsID:  sd.TimeSlotsID,
	}); err != nil {
		if errors.Is(err, repository.ErrStudentDefNotFound) {
			return ErrStudentDefNotFound
		}
		return err
	}

	return nil
}

// DeleteStudentDef deletes a member in db given by id
func (s *UserSvc) DeleteStudentDef(ctx context.Context, id int) error {
	if err := s.Repository.DeleteStudentDef(ctx, id); err != nil {
		if errors.Is(err, repository.ErrStudentDefNotFound) {
			return ErrStudentDefNotFound
		}
		return err
	}

	return nil
}

// GetStudentDefs returns a list of members in db
func (s *UserSvc) GetStudentDefs(ctx context.Context) ([]StudentDefOutputSvc, int, error) {
	sdsRepo, count, err := s.Repository.GetStudentDefs(ctx)
	if err != nil {
		return nil, 0, err
	}

	var psSvc []StudentDefOutputSvc
	for _, sd := range sdsRepo {
		psSvc = append(psSvc, StudentDefOutputSvc{
			ID:           sd.ID,
			UserID:       sd.UserID,
			InstructorID: sd.InstructorID,
			TimeSlotsID:  sd.TimeSlotsID,
		})
	}

	return psSvc, count, nil
}

// GetAllStudentDefsOfClassroom returns a list of members in a classroom in db with filter
func (s *UserSvc) GetAllStudentDefsOfInstructor(ctx context.Context, instructorID string) ([]StudentDefOutputSvc, int, error) {
	sdsRepo, count, err := s.Repository.GetAllStudentDefsOfInstructor(ctx, instructorID)
	if err != nil {
		return nil, 0, err
	}

	var usSvc []StudentDefOutputSvc
	for _, sd := range sdsRepo {
		usSvc = append(usSvc, StudentDefOutputSvc{
			ID:           sd.ID,
			UserID:       sd.UserID,
			InstructorID: sd.InstructorID,
			TimeSlotsID:  sd.TimeSlotsID,
		})
	}

	return usSvc, count, nil
}

// GetStudentDef returns a member in db given by id
func (s *UserSvc) GetStudentDefByTimeSlotsID(ctx context.Context, timeSlotsID int) (StudentDefOutputSvc, error) {
	sd, err := s.Repository.GetStudentDefByTimeSlotsID(ctx, timeSlotsID)
	if err != nil {
		if errors.Is(err, repository.ErrStudentDefNotFound) {
			return StudentDefOutputSvc{}, ErrStudentDefNotFound
		}
		return StudentDefOutputSvc{}, err
	}

	return StudentDefOutputSvc{
		ID:           sd.ID,
		UserID:       sd.UserID,
		InstructorID: sd.InstructorID,
		TimeSlotsID:  sd.TimeSlotsID,
	}, nil
}
