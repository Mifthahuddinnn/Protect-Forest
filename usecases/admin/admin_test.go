package admin

import (
	"forest/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepoAdmin struct {
	mock.Mock
}

func (m *MockRepoAdmin) RegisterAdmin(admin *entities.Admin) (*entities.Admin, error) {
	args := m.Called(admin)
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func (m *MockRepoAdmin) LoginAdmin(username, password string) (*entities.Admin, error) {
	args := m.Called(username, password)
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func (m *MockRepoAdmin) GetAdminUsername(username string) (*entities.Admin, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Admin), args.Error(1)
}

func TestUseCaseAdmin_RegisterAdmin(t *testing.T) {
	mockRepo := new(MockRepoAdmin)
	mockAdmin := &entities.Admin{Username: "test", Password: "test"}

	mockRepo.On("GetAdminUsername", mockAdmin.Username).Return(&entities.Admin{}, gorm.ErrRecordNotFound)
	mockRepo.On("RegisterAdmin", mockAdmin).Return(mockAdmin, nil)

	useCase := UseCaseAdmin{Repo: mockRepo}
	admin, err := useCase.RegisterAdmin(mockAdmin)

	assert.NoError(t, err)
	assert.NotNil(t, admin)
	mockRepo.AssertExpectations(t)
}

func TestUseCaseAdmin_LoginAdmin(t *testing.T) {
	mockRepo := new(MockRepoAdmin)
	password := "test"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	mockAdmin := &entities.Admin{Username: "test", Password: string(hashedPassword)}

	mockRepo.On("GetAdminUsername", mockAdmin.Username).Return(mockAdmin, nil)
	mockRepo.On("LoginAdmin", mockAdmin.Username, password).Return(mockAdmin, nil)

	useCase := UseCaseAdmin{Repo: mockRepo}
	admin, err := useCase.LoginAdmin(mockAdmin.Username, password)

	assert.NoError(t, err)
	assert.NotNil(t, admin)
	mockRepo.AssertExpectations(t)
}

func TestUseCaseAdmin_RegisterAdmin_EmptyFields(t *testing.T) {
	mockRepo := new(MockRepoAdmin)
	mockAdmin := &entities.Admin{Username: "", Password: ""}

	useCase := UseCaseAdmin{Repo: mockRepo}
	admin, err := useCase.RegisterAdmin(mockAdmin)

	assert.Error(t, err)
	assert.Nil(t, admin)
}

func TestUseCaseAdmin_RegisterAdmin_ExistingUsername(t *testing.T) {
	mockRepo := new(MockRepoAdmin)
	mockAdmin := &entities.Admin{Username: "test", Password: "test"}

	mockRepo.On("GetAdminUsername", mockAdmin.Username).Return(mockAdmin, nil)

	useCase := UseCaseAdmin{Repo: mockRepo}
	admin, err := useCase.RegisterAdmin(mockAdmin)

	assert.Error(t, err)
	assert.Nil(t, admin)
}

func TestUseCaseAdmin_LoginAdmin_EmptyFields(t *testing.T) {
	mockRepo := new(MockRepoAdmin)

	useCase := UseCaseAdmin{Repo: mockRepo}
	admin, err := useCase.LoginAdmin("", "")

	assert.Error(t, err)
	assert.Nil(t, admin)
}

func TestUseCaseAdmin_LoginAdmin_UsernameNotFound(t *testing.T) {
	mockRepo := new(MockRepoAdmin)
	username := "test"
	password := "test"

	mockRepo.On("GetAdminUsername", username).Return(nil, gorm.ErrRecordNotFound)

	useCase := UseCaseAdmin{Repo: mockRepo}
	admin, err := useCase.LoginAdmin(username, password)

	assert.Error(t, err)
	assert.Nil(t, admin)
}

func TestUseCaseAdmin_LoginAdmin_IncorrectPassword(t *testing.T) {
	mockRepo := new(MockRepoAdmin)
	username := "test"
	password := "test"
	incorrectPassword := "wrong"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	mockAdmin := &entities.Admin{Username: username, Password: string(hashedPassword)}

	mockRepo.On("GetAdminUsername", username).Return(mockAdmin, nil)

	useCase := UseCaseAdmin{Repo: mockRepo}
	admin, err := useCase.LoginAdmin(username, incorrectPassword)

	assert.Error(t, err)
	assert.Nil(t, admin)
}
