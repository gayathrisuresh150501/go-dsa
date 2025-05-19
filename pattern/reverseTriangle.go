package pattern

import (
	"fmt"
	"strings"
)

func PrintReverseTriangle(l int) {
	var buf strings.Builder

	for row := l; row >= 1; row -= 1 {
		for space := l - row; space >= 1; space -= 1 {
			buf.WriteString(" ")
		}

		for col := 2*row - 1; col >= 1; col -= 1 {
			buf.WriteString("*")
		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

// *****
//  ***
//   *
