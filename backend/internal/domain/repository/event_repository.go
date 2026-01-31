package repository

 import "backend/internal/domain/model"
 
 // EventRepository: イベント自体の管理
 type EventRepository interface {
	Create(event *model.Event) error
	FindByID(id uint) (*model.Event, error)
	FindAll() ([]*model.Event, error)
 }

 // AttendanceRepository: 出欠状況の管理
 type AttendanceRepository interface {
	Save(attendance *model.EventAttendance) error
 }

