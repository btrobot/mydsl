package fetch

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFetcher_Fetch(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查请求头
		userAgent := r.Header.Get("User-Agent")
		if userAgent != "TestAgent" {
			t.Errorf("User-Agent header wrong. got=%q, want=%q", userAgent, "TestAgent")
		}
		
		// 返回测试响应
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body>Test Response</body></html>"))
	}))
	defer server.Close()
	
	// 创建抓取器
	options := DefaultOptions()
	options.UserAgent = "TestAgent"
	fetcher := NewFetcher(options)
	
	// 抓取测试 URL
	ctx := context.Background()
	resp, err := fetcher.Fetch(ctx, server.URL)
	
	// 检查结果
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode wrong. got=%d, want=%d", resp.StatusCode, http.StatusOK)
	}
	
	expectedBody := "<html><body>Test Response</body></html>"
	if string(resp.Body) != expectedBody {
		t.Errorf("Body wrong. got=%q, want=%q", string(resp.Body), expectedBody)
	}
	
	if resp.URL != server.URL {
		t.Errorf("URL wrong. got=%q, want=%q", resp.URL, server.URL)
	}
}

func TestFetcher_FetchBatch(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 返回测试响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test Response"))
	}))
	defer server.Close()
	
	// 创建抓取器
	fetcher := NewFetcher(DefaultOptions())
	
	// 准备测试 URL
	urls := []string{
		server.URL + "/1",
		server.URL + "/2",
		server.URL + "/3",
	}
	
	// 批量抓取
	ctx := context.Background()
	results := fetcher.FetchBatch(ctx, urls, 2)
	
	// 收集结果
	responses := []*Response{}
	for resp := range results {
		responses = append(responses, resp)
	}
	
	// 检查结果数量
	if len(responses) != len(urls) {
		t.Fatalf("Got %d responses, want %d", len(responses), len(urls))
	}
	
	// 检查每个响应
	for _, resp := range responses {
		if resp.Error != nil {
			t.Errorf("Response has error: %v", resp.Error)
		}
		
		if resp.StatusCode != http.StatusOK {
			t.Errorf("StatusCode wrong. got=%d, want=%d", resp.StatusCode, http.StatusOK)
		}
		
		if string(resp.Body) != "Test Response" {
			t.Errorf("Body wrong. got=%q, want=%q", string(resp.Body), "Test Response")
		}
	}
}

func TestFetcher_Timeout(t *testing.T) {
	// 创建延迟响应的测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 延迟 2 秒
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Delayed Response"))
	}))
	defer server.Close()
	
	// 创建超时设置为 1 秒的抓取器
	options := DefaultOptions()
	options.Timeout = 1 * time.Second
	fetcher := NewFetcher(options)
	
	// 抓取测试 URL
	ctx := context.Background()
	_, err := fetcher.Fetch(ctx, server.URL)
	
	// 检查是否超时
	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
	
	// 检查错误类型
	if !strings.Contains(err.Error(), "timeout") && !strings.Contains(err.Error(), "deadline") {
		t.Errorf("Expected timeout error, got: %v", err)
	}
}
