MD to HTML Converter

A simple Markdown to HTML converter written in Go.
It takes a .md file as input and generates a .html file as output.

🚀 Usage

# Run directly with Go

go run main.go input.md

# Example

go run main.go sample.md

# → Generates sample.html

Now you can open the generated .html file in your browser.

🧩 Logic

# → converts to an <h1> tag

No prefix symbol → converts to a <p> (paragraph)

_text_ → converts to <em>text</em> (italic)

**text** → converts to <strong>text</strong> (bold)

📂 Example

Input (sample.md):

# Converter

This is a simple paragraph.  
It supports **bold text** and _italic text_.

# Features

- Easy to use
- Written in Go
- Converts `.md` → `.html`

📖 How It Works

The converter works in three stages:

Lexer → Scans raw Markdown text and produces tokens (#, \*_, _, text, etc.).

Parser → Builds an Abstract Syntax Tree (AST) from tokens.

Renderer → Walks the AST and outputs corresponding HTML.
