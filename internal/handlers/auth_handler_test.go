package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"replog/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockAuthService struct {
	RegisterFunc func(req *models.RegisterRequest) error
	LoginFunc    func(req *models.LoginRequest) (string, error)
}

func (m *MockAuthService) Register(req *models.RegisterRequest) error {
	return m.RegisterFunc(req)
}

func (m *MockAuthService) Login(req *models.LoginRequest) (string, error) {
	return m.LoginFunc(req)
}

// func TestRegister_Success(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockServce := &MockAuthService{
// 		RegisterFunc: func(req *models.RegisterRequest) error {
// 			return nil
// 		},
// 	}

// 	handler := NewAuthHandler(mockServce)
// 	router := gin.Default()
// 	router.POST("/register", handler.RegisterUser)
// 	body := `{"email":"test@test.com","password":"123456"}`
// 	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// }

func TestRegister(t *testing.T) {
	tests := []struct {
		name           string
		input          models.RegisterRequest
		ServiceError   error
		expectedStatus int
	}{

		{
			name:           "missing fields ",
			input:          models.RegisterRequest{Email: "", Password: ""},
			ServiceError:   nil,
			expectedStatus: 400,
		}, {

			name:           "email already exists",
			input:          models.RegisterRequest{Email: "deep@gmail.com", Password: "123456"},
			ServiceError:   errors.New("email exists"),
			expectedStatus: 409,
		}, {

			name:           "success",
			input:          models.RegisterRequest{Email: "deep@gmail.com", Password: "1234567"},
			ServiceError:   nil,
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockService := &MockAuthService{
				RegisterFunc: func(req *models.RegisterRequest) error {
					return tt.ServiceError
				},
			}

			handler := NewAuthHandler(mockService)
			router := gin.Default()
			router.POST("/register", handler.RegisterUser)

			bodyBytes, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(string(bodyBytes)))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected this")
			}

		})
	}

}
func TestLogin(t *testing.T) {
	tests := []struct {
		name           string
		inputs         models.LoginRequest
		ServiceError   error
		token          string
		ExpectedStatus int
	}{
		{name: "wrong password", inputs: models.LoginRequest{Email: "deep@gmail.com", Password: "32323232"}, ServiceError: errors.New("wrong password"), token: "", ExpectedStatus: 401},
		{name: "email not found", inputs: models.LoginRequest{Email: "deep@gmail.com", Password: "32323232"}, ServiceError: errors.New("user not exist"), token: "", ExpectedStatus: 401},
		{name: "success", inputs: models.LoginRequest{Email: "deep@gmail.com", Password: "32323232"}, ServiceError: nil, token: "token1234", ExpectedStatus: 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &MockAuthService{

				LoginFunc: func(req *models.LoginRequest) (string, error) {
					return tt.token, tt.ServiceError
				},
			}

			handler := NewAuthHandler(mockService)
			router := gin.Default()
			router.POST("/login", handler.LoginUser)

			bodyBytes, _ := json.Marshal(tt.inputs)
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(string(bodyBytes)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			if w.Code != tt.ExpectedStatus {
				t.Errorf("expecting %d got %d", tt.ExpectedStatus, w.Code)
			}

		})
	}

}
