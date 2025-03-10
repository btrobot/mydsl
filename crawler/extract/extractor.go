package extract

import (
	"bytes"
	"strings"
	
	"golang.org/x/net/html"
)

// Result 表示提取结果
type Result struct {
	Text     string
	HTML     string
	Attr     map[string]string
	Children []*Result
}

// Extractor 表示数据提取器
type Extractor struct {
	// 配置选项
}

// NewExtractor 创建新的提取器
func NewExtractor() *Extractor {
	return &Extractor{}
}

// Extract 从 HTML 中提取数据
func (e *Extractor) Extract(htmlContent string, selector string) ([]*Result, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}
	
	// 解析选择器
	sel := parseSelector(selector)
	
	// 查找匹配的节点
	nodes := findNodes(doc, sel)
	
	// 转换为结果
	results := make([]*Result, 0, len(nodes))
	for _, node := range nodes {
		results = append(results, nodeToResult(node))
	}
	
	return results, nil
}

// 解析选择器（简化版）
func parseSelector(selector string) []string {
	return strings.Split(selector, " ")
}

// 查找匹配的节点（简化版）
func findNodes(n *html.Node, sel []string) []*html.Node {
	if len(sel) == 0 {
		return []*html.Node{n}
	}
	
	var matches []*html.Node
	
	// 简化版的选择器匹配逻辑
	// 实际实现需要更复杂的选择器解析和匹配
	
	// 这里仅作为示例，实际需要实现完整的 CSS 选择器
	if n.Type == html.ElementNode && n.Data == sel[0] {
		if len(sel) == 1 {
			matches = append(matches, n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				matches = append(matches, findNodes(c, sel[1:])...)
			}
		}
	}
	
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		matches = append(matches, findNodes(c, sel)...)
	}
	
	return matches
}

// 将 HTML 节点转换为结果
func nodeToResult(n *html.Node) *Result {
	result := &Result{
		Attr:     make(map[string]string),
		Children: make([]*Result, 0),
	}
	
	// 提取文本
	if n.Type == html.TextNode {
		result.Text = n.Data
	}
	
	// 提取属性
	for _, attr := range n.Attr {
		result.Attr[attr.Key] = attr.Val
	}
	
	// 提取 HTML
	var buf bytes.Buffer
	html.Render(&buf, n)
	result.HTML = buf.String()
	
	// 提取子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result.Children = append(result.Children, nodeToResult(c))
	}
	
	return result
}
