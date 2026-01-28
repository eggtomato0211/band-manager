package usecase

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
)

// UserUsecase インターフェース
// Controllerはこのインターフェースを通してビジネスロジックを呼び出します。
type UserUsecase interface {
	RegisterUser(name string, email string) error
	GetAllUsers() ([]*model.User, error)
}

// userUsecase 構造体
// Repository（インターフェース）に依存しています。
type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase コンストラクタ
// ここで Repository の実体を受け取ります（依存性の注入）。
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

// RegisterUser 実装: ユーザー登録ロジック
func (u *userUsecase) RegisterUser(name string, email string) error {
	// ここにバリデーション（名前が空じゃないか等）を書くこともあります

	// 保存するデータの組み立て
	newUser := &model.User{
		Name:  name,
		Email: email,
		// IDはDB側で自動生成されるので指定不要
	}

	// Repositoryを呼んで保存
	return u.userRepo.Create(newUser)
}

// GetAllUsers 実装: 全員取得ロジック
func (u *userUsecase) GetAllUsers() ([]*model.User, error) {
	// そのままRepositoryを呼んでデータを返す
	return u.userRepo.FindAll()
}
