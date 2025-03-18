package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SahilBheke25/quick-farm-backend/internal/app/user/mocks"
	"github.com/SahilBheke25/quick-farm-backend/internal/models"
	"github.com/SahilBheke25/quick-farm-backend/internal/pkg/apperrors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserApiTestSuite struct {
	suite.Suite
	handler     Handler
	userService *mocks.Service
}

func (suite *UserApiTestSuite) SetupTest() {
	suite.userService = &mocks.Service{}
	suite.handler = NewHandler(suite.userService)
}

func TestUserAPITestSuite(t *testing.T) {
	suite.Run(t, new(UserApiTestSuite))
}

func (suite *UserApiTestSuite) TearDownTest() {
	suite.userService.AssertExpectations(suite.T())
}

func (s *UserApiTestSuite) TestLoginHandler() {
	testCases := []struct {
		name               string
		username           string
		password           string
		setup              func()
		expectedStatusCode int
	}{
		{
			name:     "Success",
			username: "SahilBheke",
			password: "Aim@1045",
			setup: func() {
				s.userService.On("Authenticate", mock.Anything, "SahilBheke", "Aim@1045").Return(models.User{
					Id:         1,
					Username:   "SahilBheke",
					Password:   "Aim@1045",
					First_name: "sahil",
					Last_name:  "bheke",
					Email:      "sahilbheke@gmail.com",
					Phone:      "1234567891",
					Address:    "lkajsfkljsadlf",
					Pincode:    324234,
					Uid:        123456789123,
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:     "Invaild Credentials",
			username: "SahilBheke",
			password: "Aim@1045",
			setup: func() {
				s.userService.On("Authenticate", mock.Anything, "SahilBheke", "Aim@1045").Return(models.User{}, apperrors.ErrInvalidCredentials)
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:     "Server Error",
			username: "SahilBheke",
			password: "Aim@1045",
			setup: func() {
				s.userService.On("Authenticate", mock.Anything, "SahilBheke", "Aim@1045").Return(models.User{}, apperrors.ErrDbServer)
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range testCases {
		s.SetupTest()
		s.Run(test.name, func() {
			test.setup()

			req, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"SahilBheke","password":"Aim@1045"}`))
			if err != nil {
				s.T().Fatal(err)
			}
			rr := httptest.NewRecorder()
			s.handler.Login(rr, req)

			if rr.Code != test.expectedStatusCode {
				s.T().Fatalf("expected status code: %d, got: %d", test.expectedStatusCode, rr.Code)
			}
		})
		s.TearDownTest()
	}
}
