// もう一つ例を確認したい

package sub

import "fmt"

// 2つの構造体の共通interfaceを定義
// move関数を2つの構造体に対して実装している
type Species interface {
	move() string
}

// 2種類の構造体
type Bird struct {
	Name string
}
type Human struct {
	Name string
}

// それぞれのmove関数の実装
func (b *Bird) move() string {
	return "Flying"
}
func (h *Human) move() string {
	return "Walking"
}

// メイン処理
func Main2() {
	bird := Bird{"Chirper"}
	human := Human{"Human"}

	fmt.Println(bird.move())  // Flying
	fmt.Println(human.move()) // Walking
}
