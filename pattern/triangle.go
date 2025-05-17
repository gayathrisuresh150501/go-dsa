package pattern

import (
	"fmt"
	"strings"
)

func PrintTriangle(height int) {
	var buf strings.Builder

	for row := 1; row <= height; row += 1 {
		for col := 1; col <= row; col += 1 {
			buf.WriteString("* ")
		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}
