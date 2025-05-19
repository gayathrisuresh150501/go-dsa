package pattern

import (
	"fmt"
	"strings"
)

func PrintHalfRhombus(n int) {
	var buf strings.Builder

	for row := 1; row <= n; row += 1 {
		for col := 1; col <= row; col += 1 {
			buf.WriteString("*")
		}

		buf.WriteString("\n")
	}

	for row := n - 1; row >= 1; row -= 1 {
		for col := row; col >= 1; col -= 1 {
			buf.WriteString("*")
		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

// *
// **
// ***
// **
// *
