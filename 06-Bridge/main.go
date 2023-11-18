/*
	In this design pattern we divide huge classes into smaller to handle them independently.
*/

package main

import "fmt"

type Computer interface {
	Print()
	SetPrinter()
}

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("EPSON Printer Used")
}

type HP struct{}

func (h *HP) PrintFile() {
	fmt.Println("HP Printer Used")
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request from windows.")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request from mac.")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

func main() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	windowsComputer := &Windows{}

	macComputer.SetPrinter(hpPrinter)
	windowsComputer.SetPrinter(epsonPrinter)

	macComputer.Print()

	windowsComputer.Print()
}
