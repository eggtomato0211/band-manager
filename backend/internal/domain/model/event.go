package model

import "time"

// Event イベント情報 (飲み会、ライブ、会議など)
type Event struct {
	ID   uint      `gorm:"primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name"` // "4月新歓打ち上げ"
	Date time.Time `gorm:"not null" json:"date"` // 開催日

	// イベントに紐づく出欠リスト
	Attendances []EventAttendance `gorm:"foreignKey:EventID" json:"attendances"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EventAttendance 中間テーブル: 誰がどのイベントに参加するか
type EventAttendance struct {
	EventID uint   `gorm:"primaryKey" json:"event_id"`
	UserID  string `gorm:"primaryKey;type:uuid" json:"user_id"`

	// 出欠情報とコメント
	Status  int    `gorm:"not null;default:0" json:"status"` // 0:未回答, 1:参加, 2:不参加, 3:保留
	Comment string `json:"comment"`                          // "遅れて行きます" 等

	// Join用: レスポンスにユーザー名を含めるために定義
	User User `json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}