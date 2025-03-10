package fetch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Response 表示 HTTP 响应
type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string][]string
	URL        string
	Error      error
}

// Options 表示抓取选项
type Options struct {
	Timeout       time.Duration
	UserAgent     string
	FollowRedirect bool
	MaxRetries    int
	Headers       map[string]string
}

// DefaultOptions 返回默认选项
func DefaultOptions() Options {
	return Options{
		Timeout:       10 * time.Second,
		UserAgent:     "MyDSL Crawler/1.0",
		FollowRedirect: true,
		MaxRetries:    3,
		Headers:       make(map[string]string),
	}
}

// Fetcher 表示网页抓取器
type Fetcher struct {
	client  *http.Client
	options Options
}

// NewFetcher 创建新的抓取器
func NewFetcher(options Options) *Fetcher {
	client := &http.Client{
		Timeout: options.Timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if !options.FollowRedirect {
				return http.ErrUseLastResponse
			}
			if len(via) >= 10 {
				return fmt.Errorf("too many redirects")
			}
			return nil
		},
	}
	
	return &Fetcher{
		client:  client,
		options: options,
	}
}

// Fetch 抓取单个 URL
func (f *Fetcher) Fetch(ctx context.Context, url string) (*Response, error) {
	var resp *http.Response
	var err error
	
	// 重试逻辑
	for i := 0; i <= f.options.MaxRetries; i++ {
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, err
		}
		
		// 设置请求头
		req.Header.Set("User-Agent", f.options.UserAgent)
		for k, v := range f.options.Headers {
			req.Header.Set(k, v)
		}
		
		resp, err = f.client.Do(req)
		if err == nil {
			break
		}
		
		// 如果是最后一次重试，返回错误
		if i == f.options.MaxRetries {
			return nil, err
		}
		
		// 等待一段时间后重试
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Duration(i+1) * time.Second):
			// 继续重试
		}
	}
	
	defer resp.Body.Close()
	
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
		URL:        url,
		Error:      nil,
	}, nil
}

// FetchBatch 批量抓取 URL
func (f *Fetcher) FetchBatch(ctx context.Context, urls []string, concurrency int) <-chan *Response {
	results := make(chan *Response)
	
	// 限制并发数的信号量
	semaphore := make(chan struct{}, concurrency)
	
	go func() {
		defer close(results)
		
		for _, url := range urls {
			select {
			case <-ctx.Done():
				return
			case semaphore <- struct{}{}: // 获取信号量
				go func(url string) {
					defer func() { <-semaphore } // 释放信号量
					
					resp, err := f.Fetch(ctx, url)
					if err != nil {
						results <- &Response{
							URL:   url,
							Error: err,
						}
						return
					}
					
					results <- resp
				}(url)
			}
		}
		
		// 等待所有 goroutine 完成
		for i := 0; i < concurrency; i++ {
			semaphore <- struct{}{}
		}
	}()
	
	return results
}
