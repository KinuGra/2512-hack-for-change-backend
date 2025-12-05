package model

// AddResult は「足し算の結果」を JSON として返すための構造体。
// Gin は構造体を自動的に JSON に変換してレスポンスとして返せる。
type AddResult struct {
	Result int `json:"result"`
}
