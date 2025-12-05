package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"2512-hack-for-change-backend/internal/model"
	"2512-hack-for-change-backend/internal/service"
)

// AddHandler は HTTP リクエストを受け取り Service を呼び出す役割。
// Handler → Service → Model の流れで処理が進む。
type AddHandler struct {
	service service.AddService // HandlerはServiceに依存（依存性注入の考え方）
}

// NewAddHandler は Handler の初期化を行う。
// 内部で AddService を作って紐づけている。
func NewAddHandler() *AddHandler {
	svc := service.NewAddService() // Serviceを生成
	return &AddHandler{service: svc}
}

// Add は /add?a=1&b=2 の HTTP リクエストを処理する関数。
// Gin のハンドラーとして登録される。
func (h *AddHandler) Add(c *gin.Context) {
	// クエリパラメータ a, b を取得（例: /add?a=3&b=4）
	aStr := c.Query("a")
	bStr := c.Query("b")

	// 文字列 → int に変換（例: "3" → 3）
	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter 'a'"})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter 'b'"})
		return
	}

	// ServiceのAddを呼び出してビジネスロジック実行
	result := h.service.Add(a, b)


	// JSON形式でレスポンスを返す
	c.JSON(http.StatusOK, model.AddResult{Result: result})
}
