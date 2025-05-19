package pattern

import (
	"fmt"
	"strings"
)

func PrintRhombus(n int) {
	var buf strings.Builder

	for row := 1; row <= n; row += 1 {
		for space := n - row; space >= 1; space -= 1 {
			buf.WriteString(" ")
		}

		for col := 1; col <= 2*row-1; col += 1 {
			buf.WriteString("*")
		}

		buf.WriteString("\n")
	}

	for row := n - 1; row >= 1; row -= 1 {
		for space := n - row; space >= 1; space -= 1 {
			buf.WriteString(" ")
		}

		for col := 2*row - 1; col >= 1; col -= 1 {
			buf.WriteString("*")
		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

//   *
//  ***
// *****
//  ***
//   *
