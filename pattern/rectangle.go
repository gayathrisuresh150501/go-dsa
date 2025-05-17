package patterns

import (
	"fmt"
	"strings"
)

func PrintRectangle(l, b int) {
	var buf strings.Builder

	for col := 1; col <= b; col += 1 {
		for row := 1; row <= l; row += 1 {
			// fmt.Print("* ")
			buf.WriteString("* ")
		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}
