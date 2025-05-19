package pattern

import (
	"fmt"
	"strings"
)

func PrintBinaryRightTriangle(n int) {
	var buf strings.Builder

	for row := 1; row <= n; row += 1 {
		for col := 1; col <= row; col += 1 {
			if row%2 != 0 && col%2 != 0 {
				buf.WriteString("1 ")
			}
			if row%2 == 0 && col%2 == 0 {
				buf.WriteString("1 ")
			}

			if row%2 != 0 && col%2 == 0 {
				buf.WriteString("0 ")
			}

			if row%2 == 0 && col%2 != 0 {
				buf.WriteString("0 ")
			}
		}
		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

// 1
// 0 1
// 1 0 1
// 0 1 0 1
// 1 0 1 0 1
