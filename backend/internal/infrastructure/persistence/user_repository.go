package persistence

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"gorm.io/gorm"
)

// userRepository 構造体
// DB接続情報(gorm.DB)を持たせます
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository コンストラクタ
// main.go で呼び出して、依存性を注入します
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

// Create 実装: ユーザーをDBにINSERT
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// FindAll 実装: 全ユーザーをSELECT
func (r *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	// SELECT * FROM users
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByID 実装: 特定のユーザーをSELECT
func (r *userRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	// SELECT * FROM users WHERE id = ?
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
