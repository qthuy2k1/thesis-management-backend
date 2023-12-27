package service

import (
	"context"
	"testing"

	repository "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository/user"
	"github.com/stretchr/testify/assert"
)

func Test_UserService_CreateUser(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserRepo struct {
		input   repository.UserInputRepo
		err     error
		expCall bool
	}

	testCases := map[string]struct {
		input        UserInputSvc
		mockUserRepo mockUserRepo
		err          error
	}{
		"success": {
			input: UserInputSvc{
				ID:       "123",
				Class:    &class,
				Major:    &major,
				Phone:    &phone,
				PhotoSrc: "example.com",
				Role:     "lecturer",
				Name:     "ABC",
				Email:    "example@gmail.com",
			},
			mockUserRepo: mockUserRepo{
				expCall: true,
				input: repository.UserInputRepo{
					ID:       "123",
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIUserRepo{}
			service := NewUserSvc(mockRepo)

			if tc.mockUserRepo.expCall {
				mockRepo.On("CreateUser", context.Background(), tc.mockUserRepo.input).Return(tc.mockUserRepo.err)
			}

			if err := service.CreateUser(context.Background(), tc.input); tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func Test_UserService_GetUser(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserRepo struct {
		userID  string
		err     error
		output  repository.UserOutputRepo
		expCall bool
	}

	testCases := map[string]struct {
		userID       string
		output       UserOutputSvc
		mockUserRepo mockUserRepo
		err          error
	}{
		"success": {
			userID: "123",
			output: UserOutputSvc{
				ID:       "123",
				Class:    &class,
				Major:    &major,
				Phone:    &phone,
				PhotoSrc: "example.com",
				Role:     "lecturer",
				Name:     "ABC",
				Email:    "example@gmail.com",
			},
			mockUserRepo: mockUserRepo{
				expCall: true,
				userID:  "123",
				output: repository.UserOutputRepo{
					ID:       "123",
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
		},
		"user not found": {
			userID: "123",
			mockUserRepo: mockUserRepo{
				expCall: true,
				userID:  "123",
				err:     repository.ErrUserNotFound,
			},
			err: ErrUserNotFound,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIUserRepo{}
			service := NewUserSvc(mockRepo)

			if tc.mockUserRepo.expCall {
				mockRepo.On("GetUser", context.Background(), tc.mockUserRepo.userID).Return(tc.mockUserRepo.output, tc.mockUserRepo.err)
			}

			user, err := service.GetUser(context.Background(), tc.userID)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}

			assert.Equal(t, tc.output, user)
		})
	}
}

func Test_UserService_DeleteUser(t *testing.T) {
	type mockUserRepo struct {
		userID  string
		err     error
		expCall bool
	}

	testCases := map[string]struct {
		userID       string
		mockUserRepo mockUserRepo
		err          error
	}{
		"success": {
			userID: "123",
			mockUserRepo: mockUserRepo{
				expCall: true,
				userID:  "123",
			},
		},
		"user not found": {
			userID: "123",
			mockUserRepo: mockUserRepo{
				expCall: true,
				userID:  "123",
				err:     repository.ErrUserNotFound,
			},
			err: ErrUserNotFound,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIUserRepo{}
			service := NewUserSvc(mockRepo)

			if tc.mockUserRepo.expCall {
				mockRepo.On("DeleteUser", context.Background(), tc.mockUserRepo.userID).Return(tc.mockUserRepo.err)
			}

			if err := service.DeleteUser(context.Background(), tc.userID); tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func Test_UserService_UpdateUser(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserRepo struct {
		input   repository.UserInputRepo
		userID  string
		err     error
		expCall bool
	}

	testCases := map[string]struct {
		input        UserInputSvc
		userID       string
		mockUserRepo mockUserRepo
		err          error
	}{
		"success": {
			userID: "123",
			input: UserInputSvc{
				Class:    &class,
				Major:    &major,
				Phone:    &phone,
				PhotoSrc: "example.com",
				Role:     "lecturer",
				Name:     "ABC",
				Email:    "example@gmail.com",
			},
			mockUserRepo: mockUserRepo{
				expCall: true,
				userID:  "123",
				input: repository.UserInputRepo{
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
		},
		"not found": {
			userID: "123",
			input: UserInputSvc{
				Class:    &class,
				Major:    &major,
				Phone:    &phone,
				PhotoSrc: "example.com",
				Role:     "lecturer",
				Name:     "ABC",
				Email:    "example@gmail.com",
			},
			mockUserRepo: mockUserRepo{
				expCall: true,
				userID:  "123",
				input: repository.UserInputRepo{
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
				err: repository.ErrUserNotFound,
			},
			err: ErrUserNotFound,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIUserRepo{}
			service := NewUserSvc(mockRepo)

			if tc.mockUserRepo.expCall {
				mockRepo.On("UpdateUser", context.Background(), tc.mockUserRepo.userID, tc.mockUserRepo.input).Return(tc.mockUserRepo.err)
			}

			if err := service.UpdateUser(context.Background(), tc.userID, tc.input); tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func Test_UserService_GetUsers(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserRepo struct {
		err     error
		count   int
		output  []repository.UserOutputRepo
		expCall bool
	}

	testCases := map[string]struct {
		output       []UserOutputSvc
		count        int
		mockUserRepo mockUserRepo
		err          error
	}{
		"success": {
			output: []UserOutputSvc{
				{
					ID:       "123",
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
				{
					ID:       "1234",
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
			count: 2,
			mockUserRepo: mockUserRepo{
				expCall: true,
				count:   2,
				output: []repository.UserOutputRepo{
					{
						ID:       "123",
						Class:    &class,
						Major:    &major,
						Phone:    &phone,
						PhotoSrc: "example.com",
						Role:     "lecturer",
						Name:     "ABC",
						Email:    "example@gmail.com",
					},
					{
						ID:       "1234",
						Class:    &class,
						Major:    &major,
						Phone:    &phone,
						PhotoSrc: "example.com",
						Role:     "lecturer",
						Name:     "ABC",
						Email:    "example@gmail.com",
					},
				},
			},
		},
		"user not found": {
			count: 0,
			mockUserRepo: mockUserRepo{
				count:   0,
				expCall: true,
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIUserRepo{}
			service := NewUserSvc(mockRepo)

			if tc.mockUserRepo.expCall {
				mockRepo.On("GetUsers", context.Background()).Return(tc.mockUserRepo.output, tc.mockUserRepo.count, tc.mockUserRepo.err)
			}

			user, count, err := service.GetUsers(context.Background())
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}

			assert.Equal(t, tc.output, user)
			assert.Equal(t, tc.count, count)
		})
	}
}
