package mysql

import (
	"fmt"
	"github.com/airondev/food-app-ddd-golang/domain/entity"
	"github.com/airondev/food-app-ddd-golang/domain/repository"
	"github.com/jinzhu/gorm"
)

type Repositories struct {
	User repository.UserRepository
	//Food repository.FoodRepository
	db *gorm.DB
}

func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {

	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", DbUser, DbPassword, DbHost, DbPort, DbName, "utf8mb4", "True", "Local")
	db, err := gorm.Open(DbDriver, dburl)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

// Close closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// Automigrate This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}).Error
}
