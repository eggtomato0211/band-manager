package model

import "time"

// User ユーザー情報
type User struct {
	// IDは予測されにくいUUIDを使用 (PostgreSQLの機能で自動生成)
	ID    string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"unique;not null" json:"email"`
	
	// ★将来の拡張用: ユーザーは複数の楽器を担当できる
	Instruments []Instrument `gorm:"many2many:user_instruments;" json:"instruments"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Instrument 楽器マスタ (例: Gt, Ba, Dr)
type Instrument struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}