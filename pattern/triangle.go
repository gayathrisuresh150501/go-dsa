package pattern

import (
	"fmt"
	"strings"
)

func PrintTriangle(l int) {
	var buf strings.Builder

	for row := 1; row <= l; row += 1 {
		for space := 1; space <= l-row; space += 1 {
			buf.WriteString(" ")
		}

		for col := 1; col <= 2*row-1; col += 1 {
			buf.WriteString("*")
		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

//   *
//  ***
// *****
