package report_test

import (
	"errors"
	"fmt"
	"forest/entities"
	"forest/usecases/report"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateReport(report *entities.Report) (*entities.Report, error) {
	args := m.Called(report)
	return args.Get(0).(*entities.Report), args.Error(1)
}

func (m *MockRepo) RegisterUser(user *entities.User) (*entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockRepo) LoginUser(email, password string) (*entities.User, error) {
	args := m.Called(email, password)
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

func (m *MockRepo) GetReports() ([]*entities.Report, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Report), args.Error(1)
}

func (m *MockRepo) GetReportByID(id int) (*entities.Report, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Report), args.Error(1)
}

func (m *MockRepo) DeleteReport(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepo) UpdateReport(report *entities.Report) (*entities.Report, error) {
	args := m.Called(report)
	return args.Get(0).(*entities.Report), args.Error(1)
}

func (m *MockRepo) GetUserByID(id int) (*entities.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func TestCreateReport(t *testing.T) {
	mockRepo := new(MockRepo)
	mockReport := &entities.Report{ID: 1, UserID: 1, Title: "Test", Content: "Test Content"}

	mockRepo.On("GetUserByID", mockReport.UserID).Return(&entities.User{ID: mockReport.UserID}, nil)
	mockRepo.On("CreateReport", mockReport).Return(mockReport, nil)

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	report, err := usecase.CreateReport(mockReport)

	assert.NoError(t, err)
	assert.NotNil(t, report)
	mockRepo.AssertExpectations(t)
}

func TestCreateReport_Error(t *testing.T) {
	mockRepo := new(MockRepo)
	mockReport := &entities.Report{}

	mockRepo.On("GetUserByID", mockReport.UserID).Return(nil, errors.New("error"))

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	_, err := usecase.CreateReport(mockReport)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetReports(t *testing.T) {
	mockRepo := new(MockRepo)
	mockReports := []*entities.Report{
		{ID: 1, UserID: 1, Title: "Test1", Content: "Test Content1"},
		{ID: 2, UserID: 2, Title: "Test2", Content: "Test Content2"},
	}

	mockRepo.On("GetReports").Return(mockReports, nil)

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	reports, err := usecase.GetReports()

	assert.NoError(t, err)
	assert.Equal(t, mockReports, reports)
	mockRepo.AssertExpectations(t)
}

func TestGetReportByID(t *testing.T) {
	mockRepo := new(MockRepo)
	mockReport := &entities.Report{ID: 1, UserID: 1, Title: "Test", Content: "Test Content"}

	mockRepo.On("GetReportByID", mockReport.ID).Return(mockReport, nil)

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	report, err := usecase.GetReportByID(mockReport.ID)

	assert.NoError(t, err)
	assert.NotNil(t, report)
	mockRepo.AssertExpectations(t)
}

func TestDeleteReport(t *testing.T) {
	mockRepo := new(MockRepo)
	mockReport := &entities.Report{ID: 1, UserID: 1, Title: "Test", Content: "Test Content"}

	mockRepo.On("GetReportByID", mockReport.ID).Return(mockReport, nil)
	mockRepo.On("DeleteReport", mockReport.ID).Return(nil)

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	err := usecase.DeleteReport(mockReport.ID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApproveReport_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	reportID := 1
	adminID := 1
	mockReport := &entities.Report{}

	mockRepo.On("GetReportByID", reportID).Return(mockReport, fmt.Errorf("report not found"))

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	_, err := usecase.ApproveReport(reportID, adminID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApproveReport_AlreadyApproved(t *testing.T) {
	mockRepo := new(MockRepo)
	reportID := 1
	adminID := 1
	mockReport := &entities.Report{Status: "approved"}

	mockRepo.On("GetReportByID", reportID).Return(mockReport, nil)

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	_, err := usecase.ApproveReport(reportID, adminID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApproveReport(t *testing.T) {
	mockRepo := new(MockRepo)
	reportID := 1
	adminID := 1
	mockReport := &entities.Report{Status: "pending"}

	mockRepo.On("GetReportByID", reportID).Return(mockReport, nil)
	mockRepo.On("UpdateReport", mockReport).Return(mockReport, nil)
	mockRepo.On("AddPointsToUser", mockReport.UserID, 1).Return(nil) // Add this line

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	report, err := usecase.ApproveReport(reportID, adminID)

	assert.NoError(t, err)
	assert.NotNil(t, report)
	mockRepo.AssertExpectations(t)
}

func TestApproveReport_UpdateError(t *testing.T) {
	mockRepo := new(MockRepo)
	reportID := 1
	adminID := 1
	mockReport := &entities.Report{Status: "pending"}

	mockRepo.On("GetReportByID", reportID).Return(mockReport, nil)
	mockRepo.On("UpdateReport", mockReport).Return(mockReport, errors.New("error"))

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	_, err := usecase.ApproveReport(reportID, adminID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApproveReport_AddPointsError(t *testing.T) {
	mockRepo := new(MockRepo)
	reportID := 1
	adminID := 1
	mockReport := &entities.Report{Status: "pending"}

	mockRepo.On("GetReportByID", reportID).Return(mockReport, nil)
	mockRepo.On("UpdateReport", mockReport).Return(mockReport, nil)
	mockRepo.On("AddPointsToUser", mockReport.UserID, 1).Return(errors.New("error"))

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	_, err := usecase.ApproveReport(reportID, adminID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateReport_InvalidReport(t *testing.T) {
	mockRepo := new(MockRepo)
	mockReport := &entities.Report{ID: 1, UserID: 1, Title: "", Content: ""}

	mockRepo.On("GetUserByID", mockReport.UserID).Return(&entities.User{ID: mockReport.UserID}, nil)

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	_, err := usecase.CreateReport(mockReport)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteReport_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	mockReport := &entities.Report{ID: 1, UserID: 1, Title: "Test", Content: "Test Content"}

	mockRepo.On("GetReportByID", mockReport.ID).Return(mockReport, errors.New("error"))

	usecase := report.ReportUseCase{
		Repo:     mockRepo,
		UserRepo: mockRepo,
	}
	err := usecase.DeleteReport(mockReport.ID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
