package trebuchet

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(content), "\n")
	answer := processInput(lines)
	os.WriteFile("../answer.txt", []byte(strconv.Itoa(answer)), 0644)
}

func processInput(lines []string) int {
	sum := 0
	for _, line := range lines {
		var numbers []int
		for _, char := range line {
			if char >= '0' && char <= '9' {
				numbers = append(numbers, int(char-'0'))
			}
		}
		sum += (numbers[0] * 10) + numbers[len(numbers)-1]
	}
	return sum
}
