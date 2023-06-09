package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hola, mundo!")
	err := pdf.OutputFileAndClose("holis_mundo.pdf")
	if err != nil {
		fmt.Println(err)
	}
}
