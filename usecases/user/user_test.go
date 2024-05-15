package user_test

import (
	"forest/entities"
	"forest/usecases/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) RegisterUser(user *entities.User) (*entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockRepo) LoginUser(email, password string) (*entities.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockRepo) GetUserByID(id int) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockRepo) GetUsers() ([]*entities.User, error) {
	args := m.Called()
	return args.Get(0).([]*entities.User), args.Error(1)
}

func (m *MockRepo) GetUserByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockRepo) UpdateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepo) AddPointsToUser(userID, points int) error {
	args := m.Called(userID, points)
	return args.Error(0)
}

func (m *MockRepo) RedeemPoints(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockRepo) GetBalanceByUserID(userID int) (*entities.Balance, error) {
	args := m.Called(userID)
	return args.Get(0).(*entities.Balance), args.Error(1)
}

func (m *MockRepo) UpdateBalance(balance *entities.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *MockRepo) CreateBalance(balance *entities.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockRepo)
	mockUser := &entities.User{Email: "test@test.com", Password: "password", Name: "Test User"}

	mockRepo.On("GetUserByEmail", mockUser.Email).Return(&entities.User{}, gorm.ErrRecordNotFound)
	mockRepo.On("RegisterUser", mockUser).Return(mockUser, nil)
	mockRepo.On("CreateBalance", &entities.Balance{UserID: mockUser.ID, Amount: 0}).Return(nil)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	user, err := usecase.RegisterUser(mockUser)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestLoginUser(t *testing.T) {
	mockRepo := new(MockRepo)
	mockUser := &entities.User{Email: "test@test.com", Password: "password", Name: "Test User"}

	mockRepo.On("LoginUser", mockUser.Email, mockUser.Password).Return(mockUser, nil)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	user, err := usecase.LoginUser(mockUser.Email, mockUser.Password)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, mockUser.Email, user.Email)
	assert.Equal(t, mockUser.Password, user.Password)
	mockRepo.AssertExpectations(t)
}

func TestRedeemPoints(t *testing.T) {
	mockRepo := new(MockRepo)
	userID := 1
	mockUser := &entities.User{ID: userID, Points: 10}
	mockBalance := &entities.Balance{UserID: userID, Amount: 0}

	mockRepo.On("GetUserByID", userID).Return(mockUser, nil)
	mockRepo.On("UpdateUser", mockUser).Return(nil)
	mockRepo.On("GetBalanceByUserID", userID).Return(mockBalance, nil) // Add this line
	mockRepo.On("UpdateBalance", mockBalance).Return(nil)              // Add this line

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	err := usecase.RedeemPoints(userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRedeemPoints_UserNotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	userID := 1

	mockRepo.On("GetUserByID", userID).Return(&entities.User{ID: userID, Points: 0}, nil)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	err := usecase.RedeemPoints(userID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRedeemPoints_PointNotEnough(t *testing.T) {
	mockRepo := new(MockRepo)
	userID := 1

	mockRepo.On("GetUserByID", userID).Return(&entities.User{ID: userID, Points: 0}, nil)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	err := usecase.RedeemPoints(userID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddPointsToUser(t *testing.T) {
	mockRepo := new(MockRepo)
	userID := 1
	points := 10

	mockRepo.On("AddPointsToUser", userID, points).Return(nil)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	err := usecase.AddPointsToUser(userID, points)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_UserNotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	userID := 1

	mockRepo.On("GetUserByID", userID).Return(&entities.User{}, gorm.ErrRecordNotFound)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	_, err := usecase.GetUserByID(userID)

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestGetBalanceByUserID_UserNotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	userID := 1

	mockRepo.On("GetBalanceByUserID", userID).Return((*entities.Balance)(nil), gorm.ErrRecordNotFound)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	_, err := usecase.GetBalanceByUserID(userID)

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	mockRepo.AssertExpectations(t)
}
func TestGetNews_InvalidAPIKey(t *testing.T) {
	mockRepo := new(MockRepo)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}
	_, err := usecase.GetNews()

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockRepo)
	mockUser := &entities.User{ID: 1, Email: "test@test.com", Password: "password", Name: "Test User"}

	mockRepo.On("UpdateUser", mockUser).Return(nil)

	usecase := user.UserUseCase{
		Repo: mockRepo,
	}

	err := usecase.UpdateUser(mockUser)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
