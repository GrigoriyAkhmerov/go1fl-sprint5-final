// Package personaldata.
package personaldata

import "fmt"

// Personal structure will be embedded into other structures.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print method displays the data contained in the structure on the screen.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
