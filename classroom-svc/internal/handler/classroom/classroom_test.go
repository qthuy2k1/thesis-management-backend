package handler

import (
	"context"
	"testing"
	"time"

	classroompb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/classroom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_ClassroomHandler_CreateClassroom(t *testing.T) {
	topicTags := ""
	type mockClassroomSvc struct {
		expCall  bool
		clrInput service.ClassroomInputSvc
		err      error
	}

	testCases := map[string]struct {
		givenInput       *classroompb.CreateClassroomRequest
		mockClassroomSvc mockClassroomSvc
		expResp          *classroompb.CreateClassroomResponse
		err              error
	}{
		"create classroom successfully": {
			givenInput: &classroompb.CreateClassroomRequest{
				Classroom: &classroompb.ClassroomInput{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
			},
			mockClassroomSvc: mockClassroomSvc{
				expCall: true,
				clrInput: service.ClassroomInputSvc{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
			},
			expResp: &classroompb.CreateClassroomResponse{
				Response: &classroompb.CommonClassroomResponse{
					StatusCode: 201,
					Message:    "Created",
				},
			},
		},
		"classroom already existed": {
			givenInput: &classroompb.CreateClassroomRequest{
				Classroom: &classroompb.ClassroomInput{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
			},
			mockClassroomSvc: mockClassroomSvc{
				expCall: true,
				clrInput: service.ClassroomInputSvc{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
				err: service.ErrClassroomExisted,
			},
			err: status.Errorf(codes.AlreadyExists, "err: %v", ErrClassroomExisted),
		},
		// "missing a struct field": {
		// 	givenInput: &classroompb.CreateClassroomRequest{
		// 		Classroom: &classroompb.ClassroomInput{
		// 			// Title:       "Classroom 1",
		// 			Description:     "Des 1",
		// 			Status:          "available",
		// 			LecturerID:      "123123",
		// 			ClassCourse:     "ABC",
		// 			TopicTags:       &topicTags,
		// 			QuantityStudent: 20,
		// 		},
		// 	},
		// 	mockClassroomSvc: mockClassroomSvc{
		// 		expCall: false,
		// 	},
		// 	expResp: nil,
		// 	err:     status.Errorf(codes.Internal, "err: %v", ErrServerError),
		// },
		// "quantity student is 0": {
		// 	givenInput: &classroompb.CreateClassroomRequest{
		// 		Classroom: &classroompb.ClassroomInput{
		// 			Title:           "Classroom 1",
		// 			Description:     "Des 1",
		// 			Status:          "available",
		// 			LecturerID:      "123123",
		// 			ClassCourse:     "ABC",
		// 			TopicTags:       &topicTags,
		// 			QuantityStudent: 0,
		// 		},
		// 	},
		// 	mockClassroomSvc: mockClassroomSvc{
		// 		expCall: false,
		// 	},
		// 	expResp: nil,
		// 	err:     status.Errorf(codes.Internal, "err: %v", ErrServerError),
		// },
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockSvc := &service.MockIClassroomSvc{}

			handler := NewClassroomHdl(mockSvc)

			if tc.mockClassroomSvc.expCall {
				mockSvc.On("CreateClassroom", context.Background(), tc.mockClassroomSvc.clrInput).Return(tc.mockClassroomSvc.err)
			}

			res, err := handler.CreateClassroom(context.Background(), tc.givenInput)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, err, tc.err)
			}

			if tc.mockClassroomSvc.expCall {
				mockSvc.AssertCalled(t, "CreateClassroom", context.Background(), tc.mockClassroomSvc.clrInput)
			}
		})
	}
}

func Test_ClassroomHandler_GetClassroom(t *testing.T) {
	topicTags := ""
	myCreatedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", "2023-05-11T09:01:53.102071Z")
	assert.NoError(t, err)
	myUpdatedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", "2023-05-11T09:01:53.102071Z")
	assert.NoError(t, err)

	type mockClassroomSvc struct {
		expCall     bool
		classroomID int
		output      service.ClassroomOutputSvc
		err         error
	}

	testCases := map[string]struct {
		input            *classroompb.GetClassroomRequest
		mockClassroomSvc mockClassroomSvc
		expResp          *classroompb.GetClassroomResponse
		err              error
	}{
		"get classroom successfully": {
			input: &classroompb.GetClassroomRequest{
				Id: 1,
			},
			mockClassroomSvc: mockClassroomSvc{
				expCall:     true,
				classroomID: 1,
				output: service.ClassroomOutputSvc{
					ID:              1,
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
					CreatedAt:       myCreatedTime,
					UpdatedAt:       myUpdatedTime,
				},
			},
			expResp: &classroompb.GetClassroomResponse{
				Response: &classroompb.CommonClassroomResponse{
					StatusCode: 200,
					Message:    "OK",
				},
				Classroom: &classroompb.ClassroomResponse{
					Id:              1,
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
					CreatedAt:       timestamppb.New(myCreatedTime),
					UpdatedAt:       timestamppb.New(myUpdatedTime),
				},
			},
		},
		// "invalid classroom ID": {
		// 	input: &classroompb.GetClassroomRequest{
		// 		Id: 0,
		// 	},
		// 	mockClassroomSvc: mockClassroomSvc{
		// 		expCall:     true,
		// 		classroomID: 0,
		// 	},
		// 	err: nil,
		// },
		"classroom not found": {
			input: &classroompb.GetClassroomRequest{
				Id: 1,
			},
			mockClassroomSvc: mockClassroomSvc{
				expCall:     true,
				classroomID: 1,
				err:         service.ErrClassroomNotFound,
			},
			err: status.Errorf(codes.NotFound, "err: %v", ErrClassroomNotFound),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockSvc := &service.MockIClassroomSvc{}
			handler := NewClassroomHdl(mockSvc)

			if tc.mockClassroomSvc.expCall {
				mockSvc.On("GetClassroom", context.Background(), tc.mockClassroomSvc.classroomID).Return(tc.mockClassroomSvc.output, tc.mockClassroomSvc.err)
			}

			res, err := handler.GetClassroom(context.Background(), tc.input)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			}

			if tc.mockClassroomSvc.expCall {
				mockSvc.AssertCalled(t, "GetClassroom", context.Background(), tc.mockClassroomSvc.classroomID)
			}
		})
	}
}

func Test_ClassroomHandler_UpdateClassroom(t *testing.T) {
	topicTags := ""
	type mockClassroomCtrl struct {
		expCall        bool
		classroomID    int
		classroomInput service.ClassroomInputSvc
		err            error
	}

	testCases := map[string]struct {
		givenInput        *classroompb.UpdateClassroomRequest
		mockClassroomCtrl mockClassroomCtrl
		expResp           *classroompb.UpdateClassroomResponse
		err               error
	}{
		"update classroom successfully": {
			givenInput: &classroompb.UpdateClassroomRequest{
				Id: 1,
				Classroom: &classroompb.ClassroomInput{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
			},
			mockClassroomCtrl: mockClassroomCtrl{
				expCall: true,
				classroomInput: service.ClassroomInputSvc{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
				classroomID: 1,
			},
			expResp: &classroompb.UpdateClassroomResponse{
				Response: &classroompb.CommonClassroomResponse{
					StatusCode: 200,
					Message:    "Success",
				},
			},
		},
		"classroom not found": {
			givenInput: &classroompb.UpdateClassroomRequest{
				Id: 1,
				Classroom: &classroompb.ClassroomInput{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
			},
			mockClassroomCtrl: mockClassroomCtrl{
				expCall:     true,
				classroomID: 1,
				classroomInput: service.ClassroomInputSvc{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
				err: service.ErrClassroomNotFound,
			},
			err: status.Errorf(codes.NotFound, "err: %v", ErrClassroomNotFound),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockController := &service.MockIClassroomSvc{}
			handler := NewClassroomHdl(mockController)

			if tc.mockClassroomCtrl.expCall {
				mockController.On("UpdateClassroom", context.Background(), tc.mockClassroomCtrl.classroomID, tc.mockClassroomCtrl.classroomInput).Return(tc.mockClassroomCtrl.err)
			}
			res, err := handler.UpdateClassroom(context.Background(), tc.givenInput)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			}

			if tc.mockClassroomCtrl.expCall {
				mockController.AssertCalled(t, "UpdateClassroom", context.Background(), tc.mockClassroomCtrl.classroomID, tc.mockClassroomCtrl.classroomInput)
			}
		})
	}
}

func Test_ClassroomHandler_DeleteClassroom(t *testing.T) {
	type mockClassroomCtrl struct {
		expCall     bool
		classroomID int
		err         error
	}
	// Set up the test cases
	testCases := map[string]struct {
		classroomID       int
		mockClassroomCtrl mockClassroomCtrl
		input             *classroompb.DeleteClassroomRequest
		output            *classroompb.DeleteClassroomResponse
		err               error
	}{
		"delete classroom successfully": {
			classroomID: 1,
			mockClassroomCtrl: mockClassroomCtrl{
				expCall:     true,
				classroomID: 1,
			},
			input: &classroompb.DeleteClassroomRequest{
				Id: 1,
			},
			output: &classroompb.DeleteClassroomResponse{
				Response: &classroompb.CommonClassroomResponse{
					StatusCode: 200,
					Message:    "Success",
				},
			},
		},
		// "invalid classroom ID": {
		// 	classroomID: 0,
		// 	mockClassroomCtrl: mockClassroomCtrl{
		// 		expCall: false,
		// 	},
		// },
		"classroom not found": {
			classroomID: 100,
			mockClassroomCtrl: mockClassroomCtrl{
				expCall:     true,
				err:         service.ErrClassroomNotFound,
				classroomID: 100,
			},
			input: &classroompb.DeleteClassroomRequest{
				Id: 100,
			},
			err: status.Errorf(codes.NotFound, "err: %v", ErrClassroomNotFound),
		},
	}

	// Test each test case
	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockController := &service.MockIClassroomSvc{}
			handler := NewClassroomHdl(mockController)

			if tc.mockClassroomCtrl.expCall {
				mockController.On("DeleteClassroom", context.Background(), tc.classroomID).Return(tc.mockClassroomCtrl.err)
			}
			res, err := handler.DeleteClassroom(context.Background(), tc.input)

			assert.Equal(t, tc.output, res)

			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			}

			if tc.mockClassroomCtrl.expCall {
				mockController.AssertCalled(t, "DeleteClassroom", mock.Anything, tc.classroomID)
			}
		})
	}
}

func Test_ClassroomHandler_GetClassrooms(t *testing.T) {
	topicTags := ""
	myCreatedTime, err := time.Parse("2006-01-02 15:04:05.999999", "2023-06-02 09:08:36.046843")
	assert.NoError(t, err)
	myUpdatedTime, err := time.Parse("2006-01-02 15:04:05.999999", "2023-06-02 09:08:36.046843")
	assert.NoError(t, err)
	type mockClassroomSvc struct {
		expCall bool
		query   service.ClassroomFilterSvc
		output  []service.ClassroomOutputSvc
		count   int
		err     error
	}

	testCases := map[string]struct {
		givenFilter      *classroompb.GetClassroomsRequest
		mockClassroomSvc mockClassroomSvc
		expResp          *classroompb.GetClassroomsResponse
		err              error
	}{
		"get all classrooms successfully": {
			givenFilter: &classroompb.GetClassroomsRequest{
				Limit:       5,
				Page:        1,
				TitleSearch: "",
				SortColumn:  "id",
				SortOrder:   "asc",
			},
			mockClassroomSvc: mockClassroomSvc{
				expCall: true,
				count:   2,
				query: service.ClassroomFilterSvc{
					Limit:       5,
					Page:        1,
					TitleSearch: "",
					SortColumn:  "id",
					SortOrder:   "asc",
				},
				output: []service.ClassroomOutputSvc{
					{
						ID:              1,
						Title:           "Classroom 1",
						Description:     "Des 1",
						Status:          "available",
						LecturerID:      "123123",
						ClassCourse:     "ABC",
						TopicTags:       &topicTags,
						QuantityStudent: 20,
						CreatedAt:       myCreatedTime,
						UpdatedAt:       myUpdatedTime,
					},
					{
						ID:              2,
						Title:           "Classroom 1",
						Description:     "Des 1",
						Status:          "available",
						LecturerID:      "123123",
						ClassCourse:     "ABC",
						TopicTags:       &topicTags,
						QuantityStudent: 20,
						CreatedAt:       myCreatedTime,
						UpdatedAt:       myUpdatedTime,
					},
				},
			},
			expResp: &classroompb.GetClassroomsResponse{
				Response: &classroompb.CommonClassroomResponse{
					StatusCode: 200,
					Message:    "Success",
				},
				TotalCount: 2,
				Classrooms: []*classroompb.ClassroomResponse{
					{
						Id:              1,
						Title:           "Classroom 1",
						Description:     "Des 1",
						Status:          "available",
						LecturerID:      "123123",
						ClassCourse:     "ABC",
						TopicTags:       &topicTags,
						QuantityStudent: 20,
						CreatedAt:       timestamppb.New(myCreatedTime),
						UpdatedAt:       timestamppb.New(myUpdatedTime),
					},
					{
						Id:              2,
						Title:           "Classroom 1",
						Description:     "Des 1",
						Status:          "available",
						LecturerID:      "123123",
						ClassCourse:     "ABC",
						TopicTags:       &topicTags,
						QuantityStudent: 20,
						CreatedAt:       timestamppb.New(myCreatedTime),
						UpdatedAt:       timestamppb.New(myUpdatedTime),
					},
				},
			},
		},
		"get all classrooms with filter request successfully": {
			givenFilter: &classroompb.GetClassroomsRequest{
				Page:        2,
				Limit:       1,
				TitleSearch: "",
				SortColumn:  "id",
				SortOrder:   "asc",
			},
			mockClassroomSvc: mockClassroomSvc{
				expCall: true,
				count:   1,
				query: service.ClassroomFilterSvc{
					Page:        2,
					Limit:       1,
					TitleSearch: "",
					SortColumn:  "id",
					SortOrder:   "asc",
				},
				output: []service.ClassroomOutputSvc{
					{
						ID:              2,
						Title:           "Classroom 1",
						Description:     "Des 1",
						Status:          "available",
						LecturerID:      "123123",
						ClassCourse:     "ABC",
						TopicTags:       &topicTags,
						QuantityStudent: 20,
						CreatedAt:       myCreatedTime,
						UpdatedAt:       myUpdatedTime,
					},
				},
			},
			expResp: &classroompb.GetClassroomsResponse{
				Response: &classroompb.CommonClassroomResponse{
					StatusCode: 200,
					Message:    "Success",
				},
				TotalCount: 1,
				Classrooms: []*classroompb.ClassroomResponse{
					{
						Id:              2,
						Title:           "Classroom 1",
						Description:     "Des 1",
						Status:          "available",
						LecturerID:      "123123",
						ClassCourse:     "ABC",
						TopicTags:       &topicTags,
						QuantityStudent: 20,
						CreatedAt:       timestamppb.New(myCreatedTime),
						UpdatedAt:       timestamppb.New(myUpdatedTime),
					},
				},
			},
		},
		"classrooms not found with filter": {
			givenFilter: &classroompb.GetClassroomsRequest{
				Page:        100,
				Limit:       1,
				TitleSearch: "abc",
				SortColumn:  "id",
				SortOrder:   "asc",
			},
			mockClassroomSvc: mockClassroomSvc{
				expCall: true,
				query: service.ClassroomFilterSvc{
					Page:        100,
					Limit:       1,
					TitleSearch: "abc",
					SortColumn:  "id",
					SortOrder:   "asc",
				},
				count: 0,
			},
			expResp: &classroompb.GetClassroomsResponse{
				Response: &classroompb.CommonClassroomResponse{
					StatusCode: 200,
					Message:    "Success",
				},
				TotalCount: 0,
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockController := &service.MockIClassroomSvc{}
			handler := NewClassroomHdl(mockController)

			if tc.mockClassroomSvc.expCall {
				mockController.On("GetClassrooms", context.Background(), tc.mockClassroomSvc.query).Return(tc.mockClassroomSvc.output, tc.mockClassroomSvc.count, tc.mockClassroomSvc.err)
			}
			res, err := handler.GetClassrooms(context.Background(), tc.givenFilter)

			assert.Equal(t, tc.expResp, res)

			if tc.err != nil {
				assert.Equal(t, err, tc.err)
			}

			if tc.mockClassroomSvc.expCall {
				mockController.AssertCalled(t, "GetClassrooms", context.Background(), tc.mockClassroomSvc.query)
			}
		})
	}
}
