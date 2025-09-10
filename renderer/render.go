package renderer

import "main/parser"


func RenderHTML(node *parser.Node) string {
    switch node.Type {
    case parser.NODE_DOCUMENT:
        result := ""
        for _, child := range node.Children {
            result += RenderHTML(child)
        }
        return result

    case parser.NODE_HEADING:
        return "<h1>" + node.Value + "</h1>\n"

    case parser.NODE_PARAGRAPH:
        result := "<p>"
        for _, child := range node.Children {
            result += RenderHTML(child)
        }
        result += "</p>\n"
        return result

    case parser.NODE_BOLD:
        result := "<strong>"
        for _, child := range node.Children {
            result += RenderHTML(child)
        }
        result += "</strong>"
        return result

    case parser.NODE_ITALIC:
        result := "<em>"
        for _, child := range node.Children {
            result += RenderHTML(child)
        }
        result += "</em>"
        return result

    case parser.NODE_TEXT:
        return node.Value

    default:
        return ""
    }
}
