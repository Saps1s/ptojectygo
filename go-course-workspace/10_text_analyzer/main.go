package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	onlyLines := flag.Bool("l", false, "только количество строк")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("usage: go run . [-l] <file>")
	}
	path := flag.Args()[0]

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("ошибка чтения: %v", err)
	}
	text := string(data)

	lines := strings.Count(text, "\n")
	if len(text) > 0 && !strings.HasSuffix(text, "\n") {
		lines++
	}

	if *onlyLines {
		fmt.Printf("Количество строк: %d\n", lines)
		return
	}

	fmt.Printf("--- Анализ файла: %s ---\n", path)
	fmt.Printf("Количество символов: %d\n", utf8.RuneCountInString(text))
	fmt.Printf("Количество строк: %d\n", lines)
	fmt.Printf("Количество слов: %d\n", len(strings.Fields(text)))
}