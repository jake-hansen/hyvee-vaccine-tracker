package consoleprinter

import (
	"fmt"
	"github.com/jake-hansen/hyvee-vaccine-search/domain"
)

type ConsolePrinter struct {

}

func New() *ConsolePrinter {
	return &ConsolePrinter{}
}

func (p *ConsolePrinter) Deliver(pharmacy domain.Pharmacy) error {
	fmt.Printf("New vaccine appointments available at %s\n", pharmacy.Address.Line1)
	return nil
}
