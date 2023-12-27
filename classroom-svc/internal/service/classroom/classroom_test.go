package service

import (
	"context"
	"testing"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/classroom"
	"github.com/stretchr/testify/assert"
)

func Test_ClassroomService_CreateClassroom(t *testing.T) {
	topicTags := "web"

	type mockClassroomRepo struct {
		input   repository.ClassroomInputRepo
		err     error
		expCall bool
	}

	testCases := map[string]struct {
		input             ClassroomInputSvc
		mockClassroomRepo mockClassroomRepo
		err               error
	}{
		"success": {
			input: ClassroomInputSvc{
				Title:           "Classroom 1",
				Description:     "Des 1",
				Status:          "available",
				LecturerID:      "123123",
				ClassCourse:     "ABC",
				TopicTags:       &topicTags,
				QuantityStudent: 20,
			},
			mockClassroomRepo: mockClassroomRepo{
				expCall: true,
				input: repository.ClassroomInputRepo{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIClassroomRepo{}
			service := NewClassroomSvc(mockRepo)

			if tc.mockClassroomRepo.expCall {
				mockRepo.On("CreateClassroom", context.Background(), tc.mockClassroomRepo.input).Return(tc.mockClassroomRepo.err)
			}

			if err := service.CreateClassroom(context.Background(), tc.input); tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func Test_ClassroomService_GetClassroom(t *testing.T) {
	topicTags := "web"
	myCreatedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", "2023-05-11T09:01:53.102071Z")
	assert.NoError(t, err)
	myUpdatedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", "2023-05-11T09:01:53.102071Z")
	assert.NoError(t, err)
	type mockClassroomRepo struct {
		classroomID int
		err         error
		output      repository.ClassroomOutputRepo
		expCall     bool
	}

	testCases := map[string]struct {
		classroomID       int
		output            ClassroomOutputSvc
		mockClassroomRepo mockClassroomRepo
		err               error
	}{
		"success": {
			classroomID: 1,
			output: ClassroomOutputSvc{
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
			mockClassroomRepo: mockClassroomRepo{
				expCall:     true,
				classroomID: 1,
				output: repository.ClassroomOutputRepo{
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
		},
		"classroom not found": {
			classroomID: 1,
			// output:      ClassroomOutputSvc{},
			mockClassroomRepo: mockClassroomRepo{
				expCall:     true,
				classroomID: 1,
				// output:      repository.ClassroomOutputRepo{},
				err: repository.ErrClassroomNotFound,
			},
			err: ErrClassroomNotFound,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIClassroomRepo{}
			service := NewClassroomSvc(mockRepo)

			if tc.mockClassroomRepo.expCall {
				mockRepo.On("GetClassroom", context.Background(), tc.mockClassroomRepo.classroomID).Return(tc.mockClassroomRepo.output, tc.mockClassroomRepo.err)
			}

			clr, err := service.GetClassroom(context.Background(), tc.classroomID)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}

			assert.Equal(t, tc.output, clr)
		})
	}
}

func Test_ClassroomService_DeleteClassroom(t *testing.T) {
	type mockClassroomRepo struct {
		classroomID int
		err         error
		expCall     bool
	}

	testCases := map[string]struct {
		classroomID       int
		mockClassroomRepo mockClassroomRepo
		err               error
	}{
		"success": {
			classroomID: 1,
			mockClassroomRepo: mockClassroomRepo{
				expCall:     true,
				classroomID: 1,
			},
		},
		"classroom not found": {
			classroomID: 1,
			mockClassroomRepo: mockClassroomRepo{
				expCall:     true,
				classroomID: 1,
				err:         repository.ErrClassroomNotFound,
			},
			err: ErrClassroomNotFound,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIClassroomRepo{}
			service := NewClassroomSvc(mockRepo)

			if tc.mockClassroomRepo.expCall {
				mockRepo.On("DeleteClassroom", context.Background(), tc.mockClassroomRepo.classroomID).Return(tc.mockClassroomRepo.err)
			}

			if err := service.DeleteClassroom(context.Background(), tc.classroomID); tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func Test_ClassroomService_UpdateClassroom(t *testing.T) {
	topicTags := "web"
	type mockClassroomRepo struct {
		input       repository.ClassroomInputRepo
		classroomID int
		err         error
		expCall     bool
	}

	testCases := map[string]struct {
		input             ClassroomInputSvc
		classroomID       int
		mockClassroomRepo mockClassroomRepo
		err               error
	}{
		"success": {
			classroomID: 1,
			input: ClassroomInputSvc{
				Title:           "Classroom 1",
				Description:     "Des 1",
				Status:          "available",
				LecturerID:      "123123",
				ClassCourse:     "ABC",
				TopicTags:       &topicTags,
				QuantityStudent: 20,
			},
			mockClassroomRepo: mockClassroomRepo{
				expCall:     true,
				classroomID: 1,
				input: repository.ClassroomInputRepo{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
			},
		},
		"not found": {
			classroomID: 1,
			input: ClassroomInputSvc{
				Title:           "Classroom 1",
				Description:     "Des 1",
				Status:          "available",
				LecturerID:      "123123",
				ClassCourse:     "ABC",
				TopicTags:       &topicTags,
				QuantityStudent: 20,
			},
			mockClassroomRepo: mockClassroomRepo{
				expCall:     true,
				classroomID: 1,
				input: repository.ClassroomInputRepo{
					Title:           "Classroom 1",
					Description:     "Des 1",
					Status:          "available",
					LecturerID:      "123123",
					ClassCourse:     "ABC",
					TopicTags:       &topicTags,
					QuantityStudent: 20,
				},
				err: repository.ErrClassroomNotFound,
			},
			err: ErrClassroomNotFound,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIClassroomRepo{}
			service := NewClassroomSvc(mockRepo)

			if tc.mockClassroomRepo.expCall {
				mockRepo.On("UpdateClassroom", context.Background(), tc.mockClassroomRepo.classroomID, tc.mockClassroomRepo.input).Return(tc.mockClassroomRepo.err)
			}

			if err := service.UpdateClassroom(context.Background(), tc.classroomID, tc.input); tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func Test_ClassroomService_GetClassrooms(t *testing.T) {
	topicTags := "web"
	myCreatedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", "2023-05-11T09:01:53.102071Z")
	assert.NoError(t, err)
	myUpdatedTime, err := time.Parse("2006-01-02T15:04:05.999999Z", "2023-05-11T09:01:53.102071Z")
	assert.NoError(t, err)
	type mockClassroomRepo struct {
		err     error
		filter  repository.ClassroomFilterRepo
		count   int
		output  []repository.ClassroomOutputRepo
		expCall bool
	}

	testCases := map[string]struct {
		filter            ClassroomFilterSvc
		output            []ClassroomOutputSvc
		count             int
		mockClassroomRepo mockClassroomRepo
		err               error
	}{
		"success": {
			filter: ClassroomFilterSvc{
				Limit:       5,
				Page:        1,
				TitleSearch: "",
				SortColumn:  "id",
				SortOrder:   "asc",
			},
			output: []ClassroomOutputSvc{
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
			count: 2,
			mockClassroomRepo: mockClassroomRepo{
				expCall: true,
				filter: repository.ClassroomFilterRepo{
					Limit:      5,
					Page:       1,
					SortColumn: "id",
					SortOrder:  "asc",
				},
				count: 2,
				output: []repository.ClassroomOutputRepo{
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
		},
		"classroom not found": {
			filter: ClassroomFilterSvc{
				Limit:       1,
				Page:        1000,
				TitleSearch: "asdasd",
				SortColumn:  "id",
				SortOrder:   "asc",
			},
			count: 0,
			mockClassroomRepo: mockClassroomRepo{
				filter: repository.ClassroomFilterRepo{
					Limit:      1,
					Page:       1000,
					SortColumn: "id",
					SortOrder:  "asc",
				},
				count:   0,
				expCall: true,
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			mockRepo := &repository.MockIClassroomRepo{}
			service := NewClassroomSvc(mockRepo)

			if tc.mockClassroomRepo.expCall {
				mockRepo.On("GetClassrooms", context.Background(), tc.mockClassroomRepo.filter).Return(tc.mockClassroomRepo.output, tc.mockClassroomRepo.count, tc.mockClassroomRepo.err)
			}

			clr, count, err := service.GetClassrooms(context.Background(), tc.filter)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			}

			assert.Equal(t, tc.output, clr)
			assert.Equal(t, tc.count, count)
		})
	}
}
