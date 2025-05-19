package pattern

import (
	"fmt"
	"strings"
)

func PrintBinaryRightTriangle(n int) {
	var buf strings.Builder

	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			if (row+col)%2 == 0 {
				buf.WriteString("1 ")
			} else {
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
