package isp

import "fmt"

// * Printer only *
type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {
	fmt.Println(d, "printed by MyPrinter")
}

// * Combine interfaces *
type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	fmt.Println(d, "printed by Photocopier")
}

func (p Photocopier) Scan(d Document) {
	fmt.Println(d, "scanned by Photocopier")
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// Deprecated: Combine interfaces + Decorator is preferred.
func Print(m MultiFunctionDevice, d Document) {
	m.Print(d)
}

// Deprecated: Combine interfaces + Decorator is preferred.
func Scan(m MultiFunctionDevice, d Document) {
	m.Scan(d)
}

// * Combine interfaces + Decorator *
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
