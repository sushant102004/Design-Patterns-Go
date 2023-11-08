/*
	Interface Seggrigation Principle -
	Sometime adding too many functions in functions in an interface can cause problems.
	To solve this issue we can make multiple interfaces.
*/

package main

// For example there is this interface
type Printer interface {
	Print()
	Scan()
	ColorPrint()
}

// Everything is okay with this as Modern Printers can do all three tasks
type ModernPrinter struct{}

func (m *ModernPrinter) Print() {}

func (m *ModernPrinter) Scan() {}

func (m *ModernPrinter) ColorPrint() {}

// But what if we have an old printer that can only print.
type OldPrinter struct{}

func (o *OldPrinter) Print() {}

// So in order to use OldPrinter we need to implement Printer interface with all three methods and 2 methods are useless but we still need to implement.
func (o *OldPrinter) Scan()       {}
func (o *OldPrinter) ColorPrint() {}

////////////////////////////////
// Solution
////////////////////////////////

type PrinterIntf interface {
	Print()
}

type Photocopier interface {
	Scan()
	ColorPrint()
}

type VeryOldPrinter struct{}

func (v *VeryOldPrinter) Print() {} // We are implementing only PrinterIntf

type VeryModernPrinter struct{}

func (v *VeryModernPrinter) Scan()       {} //
func (v *VeryModernPrinter) ColorPrint() {} // Implemented both PrinterIntf and Photocopier interface as Modern printers can do both.
func (v *VeryModernPrinter) Print()      {} //
