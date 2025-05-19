package pattern

import (
	"fmt"
	"strings"
)

func PrintRectangle(l, b int) {
	var buf strings.Builder

	for row := 1; row <= b; row += 1 {
		for col := 1; col <= l; col += 1 {
			// fmt.Print("* ")
			buf.WriteString("* ")
		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

// * * * * * * * * * *
// * * * * * * * * * *
// * * * * * * * * * *
// * * * * * * * * * *
// * * * * * * * * * *
