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

// Graphcs interface
type Graphcs interface {
	Draw()
}

// Triangel type
type Triangel struct{}

// Draw func paint Triangel
func (t *Triangel) Draw() {
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

// Painter paint
type Painter struct {
	graphics Graphcs
}

// SetGraphics set Graphcs
func (p *Painter) SetGraphics(g Graphcs) {
	p.graphics = g
}

// Draw func paint Painter
func (p *Painter) Draw() {
	p.graphics.Draw()
}

func main() {
	p := new(Painter)
	p.SetGraphics(new(Square))
	p.Draw()
}
