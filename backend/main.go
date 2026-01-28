package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// 自分で作ったパッケージをインポート
	"backend/internal/domain/model"
	"backend/internal/infrastructure/persistence"
	"backend/internal/interface/controller"
	"backend/internal/usecase"
)

func main() {
	// ==========================================
	// 1. インフラ層の準備 (DB接続)
	// ==========================================
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

	// マイグレーション (テーブル自動作成)
	// アプリ起動時に、Goの構造体を見てSQLテーブルを作ってくれます
	db.AutoMigrate(&model.User{}, &model.Event{}, &model.EventAttendance{})

	// ==========================================
	// 2. 依存性の注入 (DI: Dependency Injection)
	// ここで「リレーのバトン」を渡していきます
	// ==========================================

	// [Repository] DB接続(db)を渡して、Repositoryを作る
	userRepo := persistence.NewUserRepository(db)

	// [UseCase] Repositoryを渡して、UseCaseを作る
	userUsecase := usecase.NewUserUsecase(userRepo)

	// [Controller] UseCaseを渡して、Controllerを作る
	userCtrl := controller.NewUserController(userUsecase)

	// ==========================================
	// 3. Webサーバーの起動 (Gin)
	// ==========================================
	r := gin.Default()

	// ヘルスチェック (生存確認用)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// ユーザー関連のルーティング
	// controllerのメソッドを登録します
	r.POST("/users", userCtrl.SignUp)   // 登録
	r.GET("/users", userCtrl.ListUsers) // 一覧取得

	// サーバー起動 (ポート8080)
	r.Run(":8080")
}
