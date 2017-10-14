package main

import (
	"errors"
	"fmt"
	"math"
)

func GetFactoryByCountryCode(lang string) (DeviceFactory, error) {
	switch lang {
	case "RU":
		return new(RuDeviceFactory), nil
	case "EN":
		return new(EnDeviceFactory), nil
	default:
		return nil, errors.New("Unsupported Country Code: " + lang)
	}
}

func main() {
	factory, e := GetFactoryByCountryCode("RU")
	if e == nil {
		m := factory.GetMouse()
		k := factory.GetKeybord()
		t := factory.GetTouchpad()

		m.Click()
		k.Print()
		k.PrintLn()
		t.Track(10, 35)
	} else {
		fmt.Println(e)
	}

}

type Mouse interface {
	Click()
	DbClick()
	Scroll(direction int)
}

type Keybord interface {
	Print()
	PrintLn()
}

type Touchpad interface {
	Track(deltaX, deltaY int)
}

type DeviceFactory interface {
	GetMouse() Mouse
	GetKeybord() Keybord
	GetTouchpad() Touchpad
}

type RuMouse struct{}

func (rm *RuMouse) Click() {
	fmt.Println("Щелчок мышью!")
}

func (rm *RuMouse) DbClick() {
	fmt.Println("Двойной щелчок мышью!")
}

func (rm *RuMouse) Scroll(direction int) {
	if direction > 0 {
		fmt.Println("Скролим верх!")
	} else if direction < 0 {
		fmt.Println("Скролим вниз")
	} else {
		fmt.Println("Не скролим!")
	}
}

type RuKeybord struct{}

func (rk *RuKeybord) Print() {
	fmt.Println("Печатаем строку")
}

func (rk *RuKeybord) PrintLn() {
	fmt.Println("Печатаем строку с переводом строки")
}

type RuTouchpad struct{}

func (rt *RuTouchpad) Track(deltaX, deltaY int) {
	s := math.Sqrt(math.Pow(float64(deltaX), 2) + math.Pow(float64(deltaY), 2))
	fmt.Printf("Передвинулось на %d пикселей\n", int(s))
}

type EnMouse struct{}

func (em *EnMouse) Click() {
	fmt.Println("Mouse click!")
}

func (em *EnMouse) DbClick() {
	fmt.Println("Double mouse click!")
}

func (em *EnMouse) Scroll(direction int) {
	if direction > 0 {
		fmt.Println("Scroll Up")
	} else if direction < 0 {
		fmt.Println("Scroll Down")
	} else {
		fmt.Println("No scrolling")
	}
}

type EnKeybord struct{}

func (ek *EnKeybord) Print() {
	fmt.Println("Print")
}

func (ek *EnKeybord) PrintLn() {
	fmt.Println("Print Line")
}

type EnTouchpad struct{}

func (et *EnTouchpad) Track(deltaX, deltaY int) {
	s := math.Sqrt(math.Pow(float64(deltaX), 2) + math.Pow(float64(deltaY), 2))
	fmt.Printf("Moved %d pixels\n", int(s))
}

type EnDeviceFactory struct{}

func (edf *EnDeviceFactory) GetMouse() Mouse {
	return new(EnMouse)
}
func (edf *EnDeviceFactory) GetKeybord() Keybord {
	return new(EnKeybord)
}
func (edf *EnDeviceFactory) GetTouchpad() Touchpad {
	return new(EnTouchpad)
}

type RuDeviceFactory struct{}

func (ruf *RuDeviceFactory) GetMouse() Mouse {
	return new(RuMouse)
}
func (ruf *RuDeviceFactory) GetKeybord() Keybord {
	return new(RuKeybord)
}
func (ruf *RuDeviceFactory) GetTouchpad() Touchpad {
	return new(RuTouchpad)
}
