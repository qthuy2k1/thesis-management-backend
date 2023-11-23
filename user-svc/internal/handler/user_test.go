package handler

import (
	"context"
	"log"
	"testing"

	userpb "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_UserHandler_CreateUser(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserSvc struct {
		expCall   bool
		userInput service.UserInputSvc
		err       error
	}

	testCases := map[string]struct {
		givenInput  *userpb.CreateUserRequest
		mockUserSvc mockUserSvc
		expResp     *userpb.CreateUserResponse
		err         error
	}{
		"create user successfully": {
			givenInput: &userpb.CreateUserRequest{
				User: &userpb.UserInput{
					Id:       "123",
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
			mockUserSvc: mockUserSvc{
				expCall: true,
				userInput: service.UserInputSvc{
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
			expResp: &userpb.CreateUserResponse{
				Response: &userpb.CommonUserResponse{
					StatusCode: 201,
					Message:    "Created",
				},
			},
		},
		"user already existed": {
			givenInput: &userpb.CreateUserRequest{
				User: &userpb.UserInput{
					Id:       "123",
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
			mockUserSvc: mockUserSvc{
				expCall: true,
				userInput: service.UserInputSvc{
					ID:       "123",
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
				err: service.ErrUserExisted,
			},
			err: status.Errorf(codes.AlreadyExists, "err: %v", ErrUserExisted),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockSvc := &service.MockIUserSvc{}

			handler := NewUserHdl(mockSvc)

			if tc.mockUserSvc.expCall {
				mockSvc.On("CreateUser", context.Background(), tc.mockUserSvc.userInput).Return(tc.mockUserSvc.err)
			}

			res, err := handler.CreateUser(context.Background(), tc.givenInput)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, err, tc.err)
			}

			if tc.mockUserSvc.expCall {
				mockSvc.AssertCalled(t, "CreateUser", context.Background(), tc.mockUserSvc.userInput)
			}
		})
	}
}

func Test_UserHandler_GetUser(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserSvc struct {
		expCall bool
		userID  string
		output  service.UserOutputSvc
		err     error
	}

	testCases := map[string]struct {
		input       *userpb.GetUserRequest
		mockUserSvc mockUserSvc
		expResp     *userpb.GetUserResponse
		err         error
	}{
		"get user successfully": {
			input: &userpb.GetUserRequest{
				Id: "123",
			},
			mockUserSvc: mockUserSvc{
				expCall: true,
				userID:  "123",
				output: service.UserOutputSvc{
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
			expResp: &userpb.GetUserResponse{
				Response: &userpb.CommonUserResponse{
					StatusCode: 200,
					Message:    "OK",
				},
				User: &userpb.UserResponse{
					Id:       "123",
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
			input: &userpb.GetUserRequest{
				Id: "123",
			},
			mockUserSvc: mockUserSvc{
				expCall: true,
				userID:  "123",
				err:     service.ErrUserNotFound,
			},
			err: status.Errorf(codes.NotFound, "err: %v", ErrUserNotFound),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockSvc := &service.MockIUserSvc{}
			handler := NewUserHdl(mockSvc)

			if tc.mockUserSvc.expCall {
				log.Println(desc)
				mockSvc.On("GetUser", context.Background(), tc.mockUserSvc.userID).Return(tc.mockUserSvc.output, tc.mockUserSvc.err)
			}

			res, err := handler.GetUser(context.Background(), tc.input)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			}

			if tc.mockUserSvc.expCall {
				mockSvc.AssertCalled(t, "GetUser", context.Background(), tc.mockUserSvc.userID)
			}
		})
	}
}

func Test_UserHandler_UpdateUser(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserCtrl struct {
		expCall   bool
		userID    string
		userInput service.UserInputSvc
		err       error
	}

	testCases := map[string]struct {
		givenInput   *userpb.UpdateUserRequest
		mockUserCtrl mockUserCtrl
		expResp      *userpb.UpdateUserResponse
		err          error
	}{
		"update user successfully": {
			givenInput: &userpb.UpdateUserRequest{
				Id: "123",
				User: &userpb.UserInput{
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
			mockUserCtrl: mockUserCtrl{
				expCall: true,
				userInput: service.UserInputSvc{
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
				userID: "123",
			},
			expResp: &userpb.UpdateUserResponse{
				Response: &userpb.CommonUserResponse{
					StatusCode: 200,
					Message:    "Success",
				},
			},
		},
		"user not found": {
			givenInput: &userpb.UpdateUserRequest{
				Id: "123",
				User: &userpb.UserInput{
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
			},
			mockUserCtrl: mockUserCtrl{
				expCall: true,
				userID:  "123",
				userInput: service.UserInputSvc{
					Class:    &class,
					Major:    &major,
					Phone:    &phone,
					PhotoSrc: "example.com",
					Role:     "lecturer",
					Name:     "ABC",
					Email:    "example@gmail.com",
				},
				err: service.ErrUserNotFound,
			},
			err: status.Errorf(codes.NotFound, "err: %v", ErrUserNotFound),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockController := &service.MockIUserSvc{}
			handler := NewUserHdl(mockController)

			if tc.mockUserCtrl.expCall {
				mockController.On("UpdateUser", context.Background(), tc.mockUserCtrl.userID, tc.mockUserCtrl.userInput).Return(tc.mockUserCtrl.err)
			}
			res, err := handler.UpdateUser(context.Background(), tc.givenInput)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			}

			if tc.mockUserCtrl.expCall {
				mockController.AssertCalled(t, "UpdateUser", context.Background(), tc.mockUserCtrl.userID, tc.mockUserCtrl.userInput)
			}
		})
	}
}

func Test_UserHandler_DeleteUser(t *testing.T) {
	type mockUserCtrl struct {
		expCall bool
		userID  string
		err     error
	}
	// Set up the test cases
	testCases := map[string]struct {
		mockUserCtrl mockUserCtrl
		input        *userpb.DeleteUserRequest
		output       *userpb.DeleteUserResponse
		err          error
	}{
		"delete user successfully": {
			mockUserCtrl: mockUserCtrl{
				expCall: true,
				userID:  "123",
			},
			input: &userpb.DeleteUserRequest{
				Id: "123",
			},
			output: &userpb.DeleteUserResponse{
				Response: &userpb.CommonUserResponse{
					StatusCode: 200,
					Message:    "Success",
				},
			},
		},
		"user not found": {
			mockUserCtrl: mockUserCtrl{
				expCall: true,
				err:     service.ErrUserNotFound,
				userID:  "100",
			},
			input: &userpb.DeleteUserRequest{
				Id: "100",
			},
			err: status.Errorf(codes.NotFound, "err: %v", ErrUserNotFound),
		},
	}

	// Test each test case
	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockController := &service.MockIUserSvc{}
			handler := NewUserHdl(mockController)

			if tc.mockUserCtrl.expCall {
				mockController.On("DeleteUser", context.Background(), tc.mockUserCtrl.userID).Return(tc.mockUserCtrl.err)
			}
			res, err := handler.DeleteUser(context.Background(), tc.input)

			assert.Equal(t, tc.output, res)

			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			}

			if tc.mockUserCtrl.expCall {
				mockController.AssertCalled(t, "DeleteUser", mock.Anything, tc.mockUserCtrl.userID)
			}
		})
	}
}

func Test_UserHandler_GetUsers(t *testing.T) {
	class := "Class 1"
	major := "CNTT"
	phone := "1231231231"

	type mockUserSvc struct {
		expCall bool
		output  []service.UserOutputSvc
		count   int
		err     error
	}

	testCases := map[string]struct {
		input       *userpb.GetUsersRequest
		mockUserSvc mockUserSvc
		expResp     *userpb.GetUsersResponse
		err         error
	}{
		"get all users successfully": {
			input: &userpb.GetUsersRequest{},
			mockUserSvc: mockUserSvc{
				expCall: true,
				count:   2,
				output: []service.UserOutputSvc{
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
			expResp: &userpb.GetUsersResponse{
				Response: &userpb.CommonUserResponse{
					StatusCode: 200,
					Message:    "Success",
				},
				TotalCount: 2,
				Users: []*userpb.UserResponse{
					{
						Id:       "123",
						Class:    &class,
						Major:    &major,
						Phone:    &phone,
						PhotoSrc: "example.com",
						Role:     "lecturer",
						Name:     "ABC",
						Email:    "example@gmail.com",
					},
					{
						Id:       "1234",
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
		"users not found": {
			input: &userpb.GetUsersRequest{},
			mockUserSvc: mockUserSvc{
				expCall: true,
				count:   0,
			},
			expResp: &userpb.GetUsersResponse{
				Response: &userpb.CommonUserResponse{
					StatusCode: 200,
					Message:    "Success",
				},
				TotalCount: 0,
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockController := &service.MockIUserSvc{}
			handler := NewUserHdl(mockController)

			if tc.mockUserSvc.expCall {
				mockController.On("GetUsers", context.Background()).Return(tc.mockUserSvc.output, tc.mockUserSvc.count, tc.mockUserSvc.err)
			}
			res, err := handler.GetUsers(context.Background(), tc.input)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, err, tc.err)
			}

			if tc.mockUserSvc.expCall {
				mockController.AssertCalled(t, "GetUsers", context.Background())
			}
		})
	}
}
