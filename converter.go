package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	outputFile, err := os.Create("codebase_table.md")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	writer.WriteString("| Language                     | files | blank % | comment % | code |\n")
	writer.WriteString("|------------------------------|-------|---------|-----------|------|\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		
		if strings.HasPrefix(line, "Language") || strings.HasPrefix(line, "-------------------------------------------------------------------------------") || line == "" {
			continue
		}

		parts := strings.Fields(line) 
		if len(parts) < 5 {
			continue 
		}

		language := strings.Join(parts[:len(parts)-4], " ")
		files := parts[len(parts)-4]
		blank := parts[len(parts)-3]
		comment := parts[len(parts)-2]
		code := parts[len(parts)-1]

		writer.WriteString(fmt.Sprintf("| %-28s |  %s  |   %s   |    %s    |  %s |\n", language, files, blank, comment, code))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	writer.Flush()

	fmt.Println("Codebase was converted.")
}
