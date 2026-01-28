package repository

import "portfolio-band-manager-backend/internal/domain/model"

// UserRepository インターフェース
// UseCase層は、このインターフェースを通じてDBを操作します。
// 具体的なDBの中身（SQLやGORM）を知らなくても良いため、テストがしやすくなります。
type UserRepository interface {
	// ユーザーを新規作成
	Create(user *model.User) error
	
	// 全ユーザーを取得 (名簿用)
	FindAll() ([]*model.User, error)
	
	// IDで検索 (後で出欠登録に使います)
	FindByID(id string) (*model.User, error)
}