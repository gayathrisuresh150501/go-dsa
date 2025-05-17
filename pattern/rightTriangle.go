package pattern

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintRightTriangle(height int) {
	var buf strings.Builder

	for row := 1; row <= height; row += 1 {
		for col := 1; col <= row; col += 1 {
			// buf.WriteString("* ")
			// buf.WriteString(strconv.Itoa(row)) // To print numbers incremented row wise
			buf.WriteString(strconv.Itoa(col)) // To print numbers incremented col wise

		}

		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}
