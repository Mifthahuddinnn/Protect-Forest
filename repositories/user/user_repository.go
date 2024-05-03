package user

import (
	"fmt"
	"forest/entities"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type RepositoryImpl struct {
	db *gorm.DB
}

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPass     = os.Getenv("DB_PASS")
	dbDatabase = os.Getenv("DB_NAME")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

func NewUserRepositoryImpl(db *gorm.DB) (Userinterface, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbDatabase)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &RepositoryImpl{db: gormDB}, nil
}

func (repo *RepositoryImpl) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result := repo.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *RepositoryImpl) GetUserByID(id int) (entities.User, error) {
	var user entities.User
	result := repo.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (repo *RepositoryImpl) UpdateUser(user *entities.User) (*entities.User, error) {
	result := repo.db.Model(&entities.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *RepositoryImpl) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (repo *RepositoryImpl) DeleteUser(id int) error {
	result := repo.db.Delete(&entities.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *RepositoryImpl) LoginUser(email, password string) (entities.User, error) {
	var user entities.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (repo *RepositoryImpl) RegisterUser(user entities.User) (entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}
	user.Password = string(hashedPassword)

	result := repo.db.Create(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}
