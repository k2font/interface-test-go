package main

import "fmt"

// 構造体 ≒ 他の言語でいうクラス、と雑に考えていい
type Circle struct {
	Radius int
}

// 構造体Circle(他の言語でいうCircleクラス)のメソッドを定義
func (c Circle) GetArea() int {
	return 3 * c.Radius * c.Radius
}

// Square構造体を定義
type Square struct {
	Height int
}

// 構造体Squareのメソッドを定義
func (s Square) GetArea() int {
	return s.Height * s.Height
}

// Figureインターフェースを定義
// GetArea関数を設定
// 型に応じて適切な関数が呼び出される
type Figure interface {
	GetArea() int
}

// Figureの面積を計算する関数を定義
func DisplayArea(f Figure) {
	fmt.Printf("面積は%vです\n", f.GetArea())
}

func main() {
	// Circle型の変数を定義
	circle := Circle{Radius: 2}
	DisplayArea(circle) // 12

	// Square型の変数を定義
	square := Square{Height: 2}
	DisplayArea(square) // 4
}
