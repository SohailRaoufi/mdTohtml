package main

import (
	"fmt"
	"main/lexer"
	"main/parser"
	"main/renderer"
	"os"
	"path/filepath"
	"strings"
)	



func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <input.md>")
        return
    }
	    inputFile := os.Args[1]

    data, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    input := string(data)

	lexer := lexer.NewLexer(input)
	tokens := lexer.Lex()

	p := parser.NewParser(tokens)
	ast := p.Parse()

	html := renderer.RenderHTML(ast)

    base := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
    outputFile := base + ".html"

	err = os.WriteFile(outputFile, []byte(html), 0644)
    if err != nil {
        fmt.Println("Error writing output file:", err)
        return
    }

    fmt.Printf("Converted %s → %s\n", inputFile, outputFile)
}
