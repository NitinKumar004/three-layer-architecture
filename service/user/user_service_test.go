package user

import (
	"go.uber.org/mock/gomock"
	User_Model "microservice/Models/user"
	"testing"
)

func TestMockstore_InsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc      string
		input     User_Model.User
		response  string
		expecterr bool
	}{
		{
			desc: "valid case",
			input: User_Model.User{
				ID:    101,
				Name:  "nitin",
				Phone: "7488204975",
				Email: "nitinraj7488204975@gmail.com",
			}, response: "Insert user successfully",
			expecterr:   false,
		},
		{
			desc: "invalid case",
			input: User_Model.User{
				ID:    101,
				Name:  "",
				Phone: "7488204975",
				Email: "nitinraj7488204975@gmail.com",
			}, response: "all fields (name, phone, email) are required",
			expecterr:   true,
		},
		{
			desc: "invalid case",
			input: User_Model.User{
				ID:    101,
				Name:  "",
				Phone: "",
				Email: "nitinraj7488204975@gmail.com",
			}, response: "all fields (name, phone, email) are required",
			expecterr:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if !tc.expecterr {
				mockStore.EXPECT().InsertUser(tc.input).Return(tc.response, nil)
			}
			msg, err := svc.InsertUser(tc.input)
			if tc.expecterr {

				if err != nil && msg != tc.response {
					t.Errorf("Expected message '%s', got '%s'", tc.response, msg)
				}

			}

		})
	}

}
func TestMockstore_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)
	testCases := []struct {
		desc      string
		mockTasks []User_Model.User
		expected  int
		expectErr bool
	}{
		{
			desc: "Tasks with valid names",
			mockTasks: []User_Model.User{
				{ID: 1, Name: " 1", Phone: "7488204975", Email: "nitin@gmail.com"},
				{ID: 2, Name: "nitin", Phone: "7488204975", Email: "nitin@gmail.com"},
				{ID: 3, Name: "", Phone: "", Email: ""},
			},
			expected:  2,
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockStore.EXPECT().GetAllUsers().Return(tc.mockTasks, nil)

			tasks, err := svc.GetAllUsers()
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
func TestMockStore_DeleteAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	mockStore.EXPECT().DeleteAllUsers().Return("user deleted successfully", nil)
	val, err := svc.DeleteAllUsers()
	if err != nil {
		t.Error("error")
	}
	if val != "user deleted successfully" {
		t.Error("Some error")
	}
}

func TestMockStore_DeleteUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc      string
		id        int
		expected  string
		expectErr bool
	}{
		{"Valid ID", 5, "deleted user", false},
		{"Invalid ID", -3, "invalid user ID", true},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if !tc.expectErr {
				mockStore.EXPECT().DeleteUserByID(tc.id).Return(tc.expected, nil)
			}

			resp, err := svc.DeleteUserByID(tc.id)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tc.expectErr && resp != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, resp)
			}
		})
	}

}
func TestMockstore_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := New(mockStore)

	testCases := []struct {
		desc     string
		id       int
		mockUser *User_Model.User

		expectErr bool
	}{
		{"Valid ID", 5, &User_Model.User{ID: 5, Name: "nitin", Email: "nitin@gmail.com", Phone: "7488204975"}, false},
		{"Invalid ID", -3, nil, true},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if !tc.expectErr {
				mockStore.EXPECT().GetUserByID(tc.id).Return(tc.mockUser, nil)
			}

			resp, err := svc.GetUserByID(tc.id)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tc.expectErr && resp.ID != tc.id {
				t.Errorf("Expected '%d', got '%d'", tc.id, resp.ID)
			}
		})
	}

}
