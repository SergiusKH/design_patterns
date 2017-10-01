/*
*  Шаблон фасад (англ. Facade) — структурный шаблон проектирования, позволяющий
*  скрыть сложность системы путём сведения всех возможных внешних вызовов к
*  одному объекту, делегирующему их соответствующим объектам системы.
 */
package main

import (
	"fmt"
)

func main() {

	computer := new(Computer)
	computer.Copy()

}

type Power struct{}

func (p *Power) On() {
	fmt.Println("Включение питания")
}

func (p *Power) Off() {
	fmt.Println("Выключение питания")
}

type DVDrom struct {
	data bool // false
}

func (dvd *DVDrom) HasData() bool {
	return dvd.data
}
func (dvd *DVDrom) load() {
	dvd.data = true
}

func (dvd *DVDrom) unload() {
	dvd.data = false
}

type HDD struct{}

func (h *HDD) CopyFromDVD(dvd *DVDrom) {
	if dvd.HasData() {
		fmt.Println("Происходит копирование данных с диска")
	} else {
		fmt.Println("Вставте диск с данными")
	}
}

type Computer struct {
	power *Power
	dvd   *DVDrom
	hdd   *HDD
}

func (c *Computer) Copy() {
	c.power = new(Power)
	c.dvd = new(DVDrom)
	c.hdd = new(HDD)

	c.power.On()
	c.dvd.load()
	//c.dvd.unload()
	c.hdd.CopyFromDVD(c.dvd)
}
