package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("sample.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineno := 1
	var issues int

	// Example: Search for hardcoded secrets (simple pattern)
	secretPattern := regexp.MustCompile(`(?i)password\s*:=|api[_-]?key\s*:=`)

	for scanner.Scan() {
		line := scanner.Text()
		if secretPattern.MatchString(line) {
			fmt.Printf("Issue at line %d: Potential hardcoded secret\n", lineno)
			issues++
		}
		lineno++
	}
	fmt.Printf("Scan complete: %d issues found\n", issues)
}
