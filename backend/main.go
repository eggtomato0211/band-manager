package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// ★ここが backend/internal/... になっています！
	"backend/internal/domain/model"
	"backend/internal/infrastructure/persistence"
	"backend/internal/interface/controller"
	"backend/internal/usecase"
)

func main() {
	// DB接続
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// マイグレーション
	db.AutoMigrate(&model.User{}, &model.Event{}, &model.EventAttendance{})

	// ==========================================
	// 2. 依存性の注入 (DI)
	// ==========================================

	// [Repository]
	userRepo := persistence.NewUserRepository(db)
	eventRepo := persistence.NewEventRepository(db)           // 追加
	attendanceRepo := persistence.NewAttendanceRepository(db) // 追加

	// [UseCase]
	userUsecase := usecase.NewUserUsecase(userRepo)
	eventUsecase := usecase.NewEventUsecase(eventRepo, attendanceRepo) // 追加

	// [Controller]
	userCtrl := controller.NewUserController(userUsecase)
	eventCtrl := controller.NewEventController(eventUsecase) // 追加

	// ==========================================
	// 3. Webサーバーの起動
	// ==========================================
	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		// 許可するオリジン (Next.jsのURL)
		AllowOrigins:     []string{"http://localhost:3000"},
		// 許可するメソッド
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 許可するヘッダー
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		// クッキーなどの情報を送ることを許可するか
		AllowCredentials: true,
		// プリフライトリクエストのキャッシュ時間
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// ユーザーAPI
	r.POST("/users", userCtrl.SignUp)
	r.GET("/users", userCtrl.ListUsers)

	// ★追加: イベント・出欠API
	r.POST("/events", eventCtrl.CreateEvent)      // イベント作成
	r.GET("/events", eventCtrl.ListEvents)        // イベント一覧取得
	r.GET("/events/:id", eventCtrl.GetEvent)      // イベント詳細(出欠状況含む)
	r.POST("/attendances", eventCtrl.AnswerAttendance) // 出欠回答

	r.Run(":8080")
}