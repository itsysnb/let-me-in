package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	for _, char := range s {
		rows[curRow] += string(char)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	var result string
	for _, row := range rows {
		result += row
	}

	return result
}

func main() {
	input := "PAYPALISHIRING"
	numRows := 3
	zigzag := convert(input, numRows)
	fmt.Println(zigzag)
}
