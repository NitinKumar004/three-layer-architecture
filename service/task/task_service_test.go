package task

import (
	"go.uber.org/mock/gomock"
	Task_Model "microservice/Models/task"
	"testing"
)

func TestInsertask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc      string
		input     Task_Model.Task
		response  string
		expectErr bool
	}{
		{
			desc: " Valid Task Insert",
			input: Task_Model.Task{
				ID:     1,
				Name:   "Write Code",
				Status: "pending",
				UserID: 101,
			},
			response:  "Inserted Successfully",
			expectErr: false,
		},
		{
			desc: " Empty Task Name",
			input: Task_Model.Task{
				ID:     2,
				Name:   "",
				Status: "pending",
				UserID: 102,
			},
			response:  "",
			expectErr: true,
		},
		{
			desc: " Empty Status",
			input: Task_Model.Task{
				ID:     3,
				Name:   "Review PR",
				Status: "",
				UserID: 103,
			},
			response:  "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if !tc.expectErr {

				mockStore.EXPECT().
					Insertask(tc.input).
					Return(tc.response, nil)
			}

			msg, err := svc.Insertask(tc.input)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error but got nil for: %s", tc.desc)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if msg != tc.response {
					t.Errorf("Expected message '%s', got '%s'", tc.response, msg)
				}
			}
		})
	}
}
func TestService_Getalltask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc      string
		mockTasks []Task_Model.Task
		expected  int
		expectErr bool
	}{
		{
			desc: "Tasks with valid names",
			mockTasks: []Task_Model.Task{
				{ID: 1, Name: "Task 1"},
				{ID: 2, Name: ""},
				{ID: 3, Name: "Task 2"},
			},
			expected:  2,
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockStore.EXPECT().Getalltask().Return(tc.mockTasks, nil)

			tasks, err := svc.Getalltask()
			if tc.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(tasks) != tc.expected {
				t.Errorf("Expected %d tasks, got %d", tc.expected, len(tasks))
			}
		})
	}
}
func TestService_Gettaskbyid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc      string
		id        int
		mockTask  *Task_Model.Task
		expectErr bool
	}{
		{"Valid ID", 1, &Task_Model.Task{ID: 1, Name: "Mock"}, false},
		{"Invalid ID", -1, nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if !tc.expectErr {
				mockStore.EXPECT().Gettaskbyid(tc.id).Return(tc.mockTask, nil)
			}

			result, err := svc.Gettaskbyid(tc.id)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error for id=%d, got nil", tc.id)
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tc.expectErr && result.ID != tc.id {
				t.Errorf("Expected ID %d, got %d", tc.id, result.ID)
			}
		})
	}
}
func TestService_Deletetask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc      string
		id        int
		mockResp  string
		expectErr bool
	}{
		{"Valid ID", 10, "Deleted", false},
		{"Invalid ID", 0, "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if !tc.expectErr {
				mockStore.EXPECT().Deletetask(tc.id).Return(tc.mockResp, nil)
			}

			resp, err := svc.Deletetask(tc.id)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tc.expectErr && resp != tc.mockResp {
				t.Errorf("Expected '%s', got '%s'", tc.mockResp, resp)
			}
		})
	}
}

func TestService_Completetask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc      string
		id        int
		mockResp  string
		expectErr bool
	}{
		{"Valid ID", 5, "Completed", false},
		{"Invalid ID", -3, "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if !tc.expectErr {
				mockStore.EXPECT().Completetask(tc.id).Return(tc.mockResp, nil)
			}

			resp, err := svc.Completetask(tc.id)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tc.expectErr && resp != tc.mockResp {
				t.Errorf("Expected '%s', got '%s'", tc.mockResp, resp)
			}
		})
	}
}
