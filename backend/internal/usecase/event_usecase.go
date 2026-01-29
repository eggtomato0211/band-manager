package usecase

import (
	"time"
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
)

type EventUsecase interface {
	CreateEvent(name string, dateStr string) error
	GetEvent(id uint) (*model.Event, error)
	RegisterAttendance(eventID uint, userID string, status int, comment string) error
}

type eventUsecase struct {
	eventRepo      repository.EventRepository
	attendanceRepo repository.AttendanceRepository
}

// コンストラクタ: 2つのリポジトリを受け取ります
func NewEventUsecase(eRepo repository.EventRepository, aRepo repository.AttendanceRepository) EventUsecase {
	return &eventUsecase{
		eventRepo:      eRepo,
		attendanceRepo: aRepo,
	}
}

func (u *eventUsecase) CreateEvent(name string, dateStr string) error {
	// 日付文字列のパース (RFC3339形式: "2006-01-02T15:04:05Z")
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return err
	}
	event := &model.Event{
		Name: name,
		Date: t,
	}
	return u.eventRepo.Create(event)
}

func (u *eventUsecase) GetEvent(id uint) (*model.Event, error) {
	return u.eventRepo.FindByID(id)
}

// RegisterAttendance 出欠登録ロジック
func (u *eventUsecase) RegisterAttendance(eventID uint, userID string, status int, comment string) error {
	attendance := &model.EventAttendance{
		EventID: eventID,
		UserID:  userID,
		Status:  status,
		Comment: comment,
	}
	// "Save" は、既にデータがあれば更新、なければ作成をしてくれます
	return u.attendanceRepo.Save(attendance)
}