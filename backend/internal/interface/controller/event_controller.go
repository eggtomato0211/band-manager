package controller

import (
	"net/http"
	"strconv"
	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	eVusecase usecase.EventUsecase
}

func NewEventController(e usecase.EventUsecase) *EventController {
	return &EventController{eVusecase: e}
}

// CreateEvent イベント作成 (POST /events)
func (ec *EventController) CreateEvent(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Date string `json:"date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ec.eVusecase.CreateEvent(req.Name, req.Date); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}

// GetEvent イベント詳細 (GET /events/:id)
func (ec *EventController) GetEvent(c *gin.Context) {
	// URLの :id を取得して数値に変換
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := ec.eVusecase.GetEvent(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": event})
}

// AnswerAttendance 出欠回答 (POST /attendances)
func (ec *EventController) AnswerAttendance(c *gin.Context) {
	var req struct {
		EventID uint   `json:"event_id"`
		UserID  string `json:"user_id"`
		Status  int    `json:"status"` // 1:参加, 2:不参加
		Comment string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ec.eVusecase.RegisterAttendance(req.EventID, req.UserID, req.Status, req.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attendance registered"})
}