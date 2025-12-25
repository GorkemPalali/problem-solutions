package longestbalancedparentheses

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solveParenthesesProblem(s string) int32 {
	stack := make([]int, 1, len(s)+1)
	stack[0] = -1 								// başlangıç tabanı

	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else { 								// s[i] == ')'
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				cur := i - stack[len(stack)-1]
				if cur > maxLen {
					maxLen = cur
				}
			}
		}
	}
	return int32(maxLen)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)
	result := solveParenthesesProblem(s)

	fmt.Fprintf(writer, "%d\n", result)
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
