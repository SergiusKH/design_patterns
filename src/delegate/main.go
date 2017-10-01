/*
* Делегирование (англ. Delegation) — основной шаблон проектирования, в котором
* объект внешне выражает некоторое поведение, но в реальности передаёт
* ответственность за выполнение этого поведения связанному объекту.
* Шаблон делегирования является фундаментальной абстракцией, на основе которой
* реализованы другие шаблоны - композиция (также называемая агрегацией),
* примеси (mixins) и аспекты (aspects).
 */

package main

import "fmt"

// A A struct type
type A struct{}

func (a *A) f() {
	fmt.Println("func f()")
}

// B type
type B struct{}

func (b *B) f() {
	a := new(A)
	a.f()
}

// Graphics interface
type Graphics interface {
	Draw()
}

// Triangle type
type Triangle struct{}

// Draw func paint Triangle
func (t *Triangle) Draw() {
	fmt.Println("Рисуем треушольник!")
}

//Square type
type Square struct{}

//Draw func paint Square
func (s *Square) Draw() {
	fmt.Println("Рисуем квадрат!")
}

// Circle type
type Circle struct{}

// Draw func paint Circle
func (c *Circle) Draw() {
	fmt.Println("Рисуем круг!")
}

// Painter type
type Painter struct {
	graphics Graphics
}

// SetGraphics set Graphcs
func (p *Painter) SetGraphics(g Graphics) {
	p.graphics = g
}

// Draw func paint shape
func (p *Painter) Draw() {
	p.graphics.Draw()
}

func main() {

	fmt.Println("==========================================Easily==================================================")

	a := new(A)
	a.f()
	b := new(B)
	b.f()

	fmt.Println("====================================Option with Interface=========================================")

	p := new(Painter)
	p.SetGraphics(new(Square))
	p.Draw()

	fmt.Println("==================================================================================================")

}
