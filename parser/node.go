package parser

type NodeType string

const (
    NODE_DOCUMENT NodeType = "DOCUMENT"
    NODE_HEADING  NodeType = "HEADING"
    NODE_PARAGRAPH NodeType = "PARAGRAPH"
    NODE_TEXT     NodeType = "TEXT"
    NODE_BOLD     NodeType = "BOLD"
    NODE_ITALIC   NodeType = "ITALIC"
)

type Node struct {
    Type     NodeType
    Value    string   
    Children []*Node  
}