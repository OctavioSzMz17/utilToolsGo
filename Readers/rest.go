package Readers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Rs(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
