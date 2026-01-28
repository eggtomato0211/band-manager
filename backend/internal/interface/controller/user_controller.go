package controller

import (
	"net/http"
	"portfolio-band-manager-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// UserController 構造体
// 具体的なロジックは知らず、「UserUsecase」というインターフェースだけを持っています。
type UserController struct {
	uUsecase usecase.UserUsecase
}

// NewUserController コンストラクタ
// ここで Usecase を受け取ります (DI: 依存性の注入)
func NewUserController(uUsecase usecase.UserUsecase) *UserController {
	return &UserController{uUsecase: uUsecase}
}

// SignUp ユーザー登録 (POST /users)
func (uc *UserController) SignUp(c *gin.Context) {
	// 1. リクエスト(JSON)を受け取る
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Usecase (ロジック) を呼び出す
	// ここで裏側がどう動くかは気にしない。「登録して！」と頼むだけ。
	if err := uc.uUsecase.RegisterUser(req.Name, req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. レスポンスを返す
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// ListUsers ユーザー一覧 (GET /users)
func (uc *UserController) ListUsers(c *gin.Context) {
	// Usecase からデータをもらう
	users, err := uc.uUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	// JSONにして返す
	c.JSON(http.StatusOK, gin.H{"users": users})
}