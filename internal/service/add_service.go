package service

// AddService はビジネスロジックのインターフェース。
// 「Add という処理があるよ」とだけ宣言し、
// 実際の実装は addService が担当する。
//
// → なぜこうする？
//   ・後で実装を差し替えられる（テスト・拡張がしやすい）
//   ・依存性が弱まり大規模開発で扱いやすい
type AddService interface {
	Add(a, b int) int
}

// addService は AddService の具体的な実装。
// struct{} は空の構造体で「中の状態を持たない」ことを示す。
type addService struct{}

// NewAddService は AddService を作るための関数（コンストラクタ的な役割）。
// Go にコンストラクタは無いのでこの書き方が慣習。
func NewAddService() AddService {
	// &addService{} は「addServiceのポインタ」を返す
	return &addService{}
}

// (s *addService) Add は addService のメソッド。
// レシーバー (s *addService) により、この関数が addService に属していると分かる。
func (s *addService) Add(a, b int) int {
	return a + b
}

