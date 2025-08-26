package parser

import (
	"fmt"
	"main/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) currentToken() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.TOKEN_EOF, Value: ""}
	}
	return p.tokens[p.pos]
}

func (p *Parser) nextToken() lexer.Token {
	tok := p.currentToken()
	p.pos++
	return tok
}

func (p *Parser) Parse() *Node {
	doc := &Node{Type: NODE_DOCUMENT, Children: []*Node{}}

	for p.currentToken().Type != lexer.TOKEN_EOF {
		node := p.parseBlock()
		if node != nil {
			doc.Children = append(doc.Children, node)
		}
	}
	return doc
}

func (p *Parser) parseBlock() *Node {
	tok := p.currentToken()

    switch tok.Type {
    case lexer.TOKEN_HEADING:
        p.nextToken()
        textTok := p.nextToken() 
        return &Node{
            Type:  NODE_HEADING,
            Value: textTok.Value,
        }

    case lexer.TOKEN_TEXT:
        return p.parseParagraph()

    case lexer.TOKEN_NEWLINE:
        p.nextToken() 
        return nil

    default:
        p.nextToken()
        return nil
    }
}

func (p *Parser) parseParagraph() *Node {
    para := &Node{Type: NODE_PARAGRAPH, Children: []*Node{}}

    for {
        tok := p.currentToken()
        if tok.Type == lexer.TOKEN_NEWLINE || tok.Type == lexer.TOKEN_EOF {
            p.nextToken()
            break
        }

        switch tok.Type {
			case lexer.TOKEN_TEXT:
				para.Children = append(para.Children, &Node{
					Type:  NODE_TEXT,
					Value: tok.Value,
				})
				p.nextToken()

			case lexer.TOKEN_BOLD_START:
				p.nextToken()
				boldNode := &Node{Type: NODE_BOLD, Children: []*Node{}}
				for p.currentToken().Type != lexer.TOKEN_BOLD_END && p.currentToken().Type != lexer.TOKEN_EOF {
					boldNode.Children = append(boldNode.Children, &Node{
						Type:  NODE_TEXT,
						Value: p.currentToken().Value,
					})
					p.nextToken()
				}
				p.nextToken() // consume '**' end
				para.Children = append(para.Children, boldNode)

			case lexer.TOKEN_ITALIC_START:
				p.nextToken() // consume '*'
				italicNode := &Node{Type: NODE_ITALIC, Children: []*Node{}}
				for p.currentToken().Type != lexer.TOKEN_ITALIC_END && p.currentToken().Type != lexer.TOKEN_EOF {
					italicNode.Children = append(italicNode.Children, &Node{
						Type:  NODE_TEXT,
						Value: p.currentToken().Value,
					})
					p.nextToken()
				}
				p.nextToken() 
				para.Children = append(para.Children, italicNode)

			default:
				p.nextToken()
        }
    }

    return para
}


func printAST(node *Node, indent string) {
    fmt.Printf("%s%s", indent, node.Type)
    if node.Value != "" {
        fmt.Printf(": %s", node.Value)
    }
    fmt.Println()
    for _, child := range node.Children {
        printAST(child, indent+"  ")
    }
}