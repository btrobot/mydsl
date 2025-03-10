package extract

import (
	"strings"
	"testing"
)

func TestExtractor_Extract(t *testing.T) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>Test Page</title>
</head>
<body>
	<h1>Hello World</h1>
	<div class="container">
		<p>First paragraph</p>
		<p>Second paragraph</p>
	</div>
	<ul>
		<li>Item 1</li>
		<li>Item 2</li>
		<li>Item 3</li>
	</ul>
</body>
</html>
`
	
	extractor := NewExtractor()
	
	// 测试提取 h1
	results, err := extractor.Extract(html, "h1")
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	
	if len(results) != 1 {
		t.Fatalf("Got %d results, want 1", len(results))
	}
	
	if !strings.Contains(results[0].Text, "Hello World") {
		t.Errorf("Text wrong. got=%q, want to contain %q", results[0].Text, "Hello World")
	}
	
	// 测试提取所有段落
	results, err = extractor.Extract(html, "p")
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	
	if len(results) != 2 {
		t.Fatalf("Got %d results, want 2", len(results))
	}
	
	expectedTexts := []string{"First paragraph", "Second paragraph"}
	for i, result := range results {
		if !strings.Contains(result.Text, expectedTexts[i]) {
			t.Errorf("Text wrong. got=%q, want to contain %q", result.Text, expectedTexts[i])
		}
	}
	
	// 测试提取列表项
	results, err = extractor.Extract(html, "li")
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	
	if len(results) != 3 {
		t.Fatalf("Got %d results, want 3", len(results))
	}
	
	expectedTexts = []string{"Item 1", "Item 2", "Item 3"}
	for i, result := range results {
		if !strings.Contains(result.Text, expectedTexts[i]) {
			t.Errorf("Text wrong. got=%q, want to contain %q", result.Text, expectedTexts[i])
		}
	}
}

func TestExtractor_ExtractWithAttributes(t *testing.T) {
	html := `
<!DOCTYPE html>
<html>
<body>
	<a href="https://example.com" id="link1">Example Link</a>
	<img src="image.jpg" alt="Example Image" width="100" height="100">
	<div class="container" data-id="123">Content</div>
</body>
</html>
`
	
	extractor := NewExtractor()
	
	// 测试提取链接
	results, err := extractor.Extract(html, "a")
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	
	if len(results) != 1 {
		t.Fatalf("Got %d results, want 1", len(results))
	}
	
	if results[0].Attr["href"] != "https://example.com" {
		t.Errorf("href attribute wrong. got=%q, want=%q", 
			results[0].Attr["href"], "https://example.com")
	}
	
	if results[0].Attr["id"] != "link1" {
		t.Errorf("id attribute wrong. got=%q, want=%q", 
			results[0].Attr["id"], "link1")
	}
	
	// 测试提取图片
	results, err = extractor.Extract(html, "img")
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	
	if len(results) != 1 {
		t.Fatalf("Got %d results, want 1", len(results))
	}
	
	if results[0].Attr["src"] != "image.jpg" {
		t.Errorf("src attribute wrong. got=%q, want=%q", 
			results[0].Attr["src"], "image.jpg")
	}
	
	if results[0].Attr["alt"] != "Example Image" {
		t.Errorf("alt attribute wrong. got=%q, want=%q", 
			results[0].Attr["alt"], "Example Image")
	}
	
	// 测试提取 div
	results, err = extractor.Extract(html, "div")
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	
	if len(results) != 1 {
		t.Fatalf("Got %d results, want 1", len(results))
	}
	
	if results[0].Attr["class"] != "container" {
		t.Errorf("class attribute wrong. got=%q, want=%q", 
			results[0].Attr["class"], "container")
	}
	
	if results[0].Attr["data-id"] != "123" {
		t.Errorf("data-id attribute wrong. got=%q, want=%q", 
			results[0].Attr["data-id"], "123")
	}
}
