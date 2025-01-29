package main

import (
	"bufio"
	"os"
)

func WriteFile(filename string, data []string) {

	file, _ := os.OpenFile(filename+".txt", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	datawriter := bufio.NewWriter(file)
	defer datawriter.Flush()

	for _, row := range data {
		_, _ = datawriter.WriteString(row + "\n")
	}
}
