package persistence

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"gorm.io/gorm"
)

// --- Event Repository ---
type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) repository.EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) Create(event *model.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) FindByID(id uint) (*model.Event, error) {
	var event model.Event
	// Preloadの連鎖: Event -> Attendances -> User
	// これで「イベント」の中に「出欠リスト」、さらにその中の「ユーザー名」まで全部取れます
	if err := r.db.Preload("Attendances.User").First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) FindAll() ([]*model.Event, error) {
	var events []*model.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

// --- Attendance Repository ---
type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) repository.AttendanceRepository {
	return &attendanceRepository{db: db}
}

// Save: 新規作成も更新もこれ1つ (Upsert)
func (r *attendanceRepository) Save(attendance *model.EventAttendance) error {
	return r.db.Save(attendance).Error
}
