package service

import (
	"context"
	"errors"

	repository "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository"
)

type UserInputSvc struct {
	ID          string
	Email       string
	Class       string
	Major       *string
	Phone       *string
	PhotoSrc    string
	Role        string
	Name        string
	ClassroomID *int
}

// CreateUser creates a new user in db given by user model
func (s *UserSvc) CreateUser(ctx context.Context, u UserInputSvc) error {
	pRepo := repository.UserInputRepo{
		ID:          u.ID,
		Class:       u.Class,
		Major:       u.Major,
		Phone:       u.Phone,
		PhotoSrc:    u.PhotoSrc,
		Role:        u.Role,
		Name:        u.Name,
		Email:       u.Email,
		ClassroomID: u.ClassroomID,
	}

	if err := s.Repository.CreateUser(ctx, pRepo); err != nil {
		if errors.Is(err, repository.ErrUserExisted) {
			return ErrUserExisted
		}
		return err
	}

	return nil
}

// GetUser returns a user in db given by id
func (s *UserSvc) GetUser(ctx context.Context, id string) (UserOutputSvc, error) {
	u, err := s.Repository.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return UserOutputSvc{}, ErrUserNotFound
		}
		return UserOutputSvc{}, err
	}

	return UserOutputSvc{
		ID:          u.ID,
		Class:       u.Class,
		Major:       u.Major,
		Phone:       u.Phone,
		PhotoSrc:    u.PhotoSrc,
		Role:        u.Role,
		Name:        u.Name,
		Email:       u.Email,
		ClassroomID: u.ClassroomID,
	}, nil
}

// UpdateUser updates the specified user by id
func (s *UserSvc) UpdateUser(ctx context.Context, id string, user UserInputSvc) error {
	if err := s.Repository.UpdateUser(ctx, id, repository.UserInputRepo{
		Class:       user.Class,
		Major:       user.Major,
		Phone:       user.Phone,
		PhotoSrc:    user.PhotoSrc,
		Role:        user.Role,
		Name:        user.Name,
		Email:       user.Email,
		ClassroomID: user.ClassroomID,
	}); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}

// DeleteUser deletes a user in db given by id
func (s *UserSvc) DeleteUser(ctx context.Context, id string) error {
	if err := s.Repository.DeleteUser(ctx, id); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}

type UserOutputSvc struct {
	ID          string
	Email       string
	Class       string
	Major       *string
	Phone       *string
	PhotoSrc    string
	Role        string
	Name        string
	ClassroomID *int
}

// GetUsers returns a list of users in db
func (s *UserSvc) GetUsers(ctx context.Context) ([]UserOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetUsers(ctx)
	if err != nil {
		return nil, 0, err
	}

	var psSvc []UserOutputSvc
	for _, u := range psRepo {
		psSvc = append(psSvc, UserOutputSvc{
			ID:          u.ID,
			Class:       u.Class,
			Major:       u.Major,
			Phone:       u.Phone,
			PhotoSrc:    u.PhotoSrc,
			Role:        u.Role,
			Name:        u.Name,
			Email:       u.Email,
			ClassroomID: u.ClassroomID,
		})
	}

	return psSvc, count, nil
}

// GetAllUsersOfClassroom returns a list of users in a classroom in db with filter
func (s *UserSvc) GetAllUsersOfClassroom(ctx context.Context, classroomID int) ([]UserOutputSvc, int, error) {
	usRepo, count, err := s.Repository.GetAllUsersOfClassroom(ctx, classroomID)
	if err != nil {
		return nil, 0, err
	}

	var usSvc []UserOutputSvc
	for _, u := range usRepo {
		usSvc = append(usSvc, UserOutputSvc{
			ID:          u.ID,
			Class:       u.Class,
			Major:       u.Major,
			Phone:       u.Phone,
			PhotoSrc:    u.PhotoSrc,
			Role:        u.Role,
			Name:        u.Name,
			Email:       u.Email,
			ClassroomID: u.ClassroomID,
		})
	}

	return usSvc, count, nil
}
