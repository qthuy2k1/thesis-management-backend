package service

import (
	"context"
	"errors"

	repository "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository/user"
)

type UserInputSvc struct {
	ID             string
	Email          string
	Class          *string
	Major          *string
	Phone          *string
	PhotoSrc       string
	Role           string
	Name           string
	HashedPassword *string
}

// CreateUser creates a new user in db given by user model
func (s *UserSvc) CreateUser(ctx context.Context, u UserInputSvc) error {
	pRepo := repository.UserInputRepo{
		ID:             u.ID,
		Class:          u.Class,
		Major:          u.Major,
		Phone:          u.Phone,
		PhotoSrc:       u.PhotoSrc,
		Role:           u.Role,
		Name:           u.Name,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
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
		ID:             u.ID,
		Class:          u.Class,
		Major:          u.Major,
		Phone:          u.Phone,
		PhotoSrc:       u.PhotoSrc,
		Role:           u.Role,
		Name:           u.Name,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
	}, nil
}

// UpdateUser updates the specified user by id
func (s *UserSvc) UpdateUser(ctx context.Context, id string, user UserInputSvc) error {
	if err := s.Repository.UpdateUser(ctx, id, repository.UserInputRepo{
		Class:          user.Class,
		Major:          user.Major,
		Phone:          user.Phone,
		PhotoSrc:       user.PhotoSrc,
		Role:           user.Role,
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
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
	ID             string
	Email          string
	Class          *string
	Major          *string
	Phone          *string
	PhotoSrc       string
	Role           string
	Name           string
	HashedPassword *string
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
			ID:             u.ID,
			Class:          u.Class,
			Major:          u.Major,
			Phone:          u.Phone,
			PhotoSrc:       u.PhotoSrc,
			Role:           u.Role,
			Name:           u.Name,
			Email:          u.Email,
			HashedPassword: u.HashedPassword,
		})
	}

	return psSvc, count, nil
}

// GetAllLecturers returns all members who has the role named "lecturer"
func (s *UserSvc) GetAllLecturers(ctx context.Context) ([]UserOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetAllLecturers(ctx)
	if err != nil {
		return nil, 0, err
	}

	var psSvc []UserOutputSvc
	for _, u := range psRepo {
		psSvc = append(psSvc, UserOutputSvc{
			ID:             u.ID,
			Class:          u.Class,
			Major:          u.Major,
			Phone:          u.Phone,
			PhotoSrc:       u.PhotoSrc,
			Role:           u.Role,
			Name:           u.Name,
			Email:          u.Email,
			HashedPassword: u.HashedPassword,
		})
	}

	return psSvc, count, nil
}

// UnsubscribeClassroom returns an error if delete occurs any errors
func (s *UserSvc) UnsubscribeClassroom(ctx context.Context, userID string, classroomID int) error {
	if err := s.Repository.UnsubscribeClassroom(ctx, userID, classroomID); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}
