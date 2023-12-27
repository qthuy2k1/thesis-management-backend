package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository/user"
)

type MemberInputSvc struct {
	ClassroomID int
	MemberID    string
	Status      string
	IsDefense   bool
}

type MemberOutputSvc struct {
	ID          int
	ClassroomID int
	MemberID    string
	Status      string
	IsDefense   bool
	CreatedAt   time.Time
}

// CreateMember creates a new member in db given by member model
func (s *UserSvc) CreateMember(ctx context.Context, m MemberInputSvc) error {
	pRepo := repository.MemberInputRepo{
		ClassroomID: m.ClassroomID,
		MemberID:    m.MemberID,
		Status:      m.Status,
		IsDefense:   m.IsDefense,
	}

	if err := s.Repository.CreateMember(ctx, pRepo); err != nil {
		if errors.Is(err, repository.ErrMemberExisted) {
			return ErrMemberExisted
		}
		return err
	}

	return nil
}

// GetMember returns a member in db given by id
func (s *UserSvc) GetMember(ctx context.Context, id int) (MemberOutputSvc, error) {
	m, err := s.Repository.GetMember(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrMemberNotFound) {
			return MemberOutputSvc{}, ErrMemberNotFound
		}
		return MemberOutputSvc{}, err
	}

	return MemberOutputSvc{
		ID:          m.ID,
		ClassroomID: m.ClassroomID,
		MemberID:    m.MemberID,
		Status:      m.Status,
		IsDefense:   m.IsDefense,
		CreatedAt:   m.CreatedAt,
	}, nil
}

// UpdateMember updates the specified member by id
func (s *UserSvc) UpdateMember(ctx context.Context, id int, member MemberInputSvc) error {
	if err := s.Repository.UpdateMember(ctx, id, repository.MemberInputRepo{
		ClassroomID: member.ClassroomID,
		MemberID:    member.MemberID,
		Status:      member.Status,
		IsDefense:   member.IsDefense,
	}); err != nil {
		if errors.Is(err, repository.ErrMemberNotFound) {
			return ErrMemberNotFound
		}
		return err
	}

	return nil
}

// DeleteMember deletes a member in db given by id
func (s *UserSvc) DeleteMember(ctx context.Context, id int) error {
	if err := s.Repository.DeleteMember(ctx, id); err != nil {
		if errors.Is(err, repository.ErrMemberNotFound) {
			return ErrMemberNotFound
		}
		return err
	}

	return nil
}

// GetMembers returns a list of members in db
func (s *UserSvc) GetMembers(ctx context.Context) ([]MemberOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetMembers(ctx)
	if err != nil {
		return nil, 0, err
	}

	var psSvc []MemberOutputSvc
	for _, m := range psRepo {
		psSvc = append(psSvc, MemberOutputSvc{
			ID:          m.ID,
			ClassroomID: m.ClassroomID,
			MemberID:    m.MemberID,
			Status:      m.Status,
			IsDefense:   m.IsDefense,
			CreatedAt:   m.CreatedAt,
		})
	}

	return psSvc, count, nil
}

// GetAllMembersOfClassroom returns a list of members in a classroom in db with filter
func (s *UserSvc) GetAllMembersOfClassroom(ctx context.Context, classroomID int) ([]MemberOutputSvc, int, error) {
	usRepo, count, err := s.Repository.GetAllMembersOfClassroom(ctx, classroomID)
	if err != nil {
		return nil, 0, err
	}

	var usSvc []MemberOutputSvc
	for _, m := range usRepo {
		usSvc = append(usSvc, MemberOutputSvc{
			ID:          m.ID,
			ClassroomID: m.ClassroomID,
			MemberID:    m.MemberID,
			Status:      m.Status,
			IsDefense:   m.IsDefense,
			CreatedAt:   m.CreatedAt,
		})
	}

	return usSvc, count, nil
}

// IsUserJoinedClassroom returns a member if exists
func (s *UserSvc) IsUserJoinedClassroom(ctx context.Context, userID string) (MemberOutputSvc, error) {
	m, err := s.Repository.IsUserJoinedClassroom(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrMemberNotFound) {
			return MemberOutputSvc{}, ErrMemberNotFound
		}
		return MemberOutputSvc{}, err
	}

	return MemberOutputSvc{
		ID:          m.ID,
		ClassroomID: m.ClassroomID,
		MemberID:    m.MemberID,
		Status:      m.Status,
		IsDefense:   m.IsDefense,
		CreatedAt:   m.CreatedAt,
	}, nil
}

// GetMember returns a member in db given by id
func (s *UserSvc) GetUserMember(ctx context.Context, userID string) (MemberOutputSvc, error) {
	m, err := s.Repository.GetUserMember(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrMemberNotFound) {
			return MemberOutputSvc{}, ErrMemberNotFound
		}
		return MemberOutputSvc{}, err
	}

	return MemberOutputSvc{
		ID:          m.ID,
		ClassroomID: m.ClassroomID,
		MemberID:    m.MemberID,
		Status:      m.Status,
		IsDefense:   m.IsDefense,
		CreatedAt:   m.CreatedAt,
	}, nil
}
