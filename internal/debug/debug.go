package debug

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	// DebugMode 控制是否启用调试输出
	DebugMode bool
	
	// DebugIndent 当前调试输出的缩进级别
	DebugIndent int
	
	// DebugWriter 调试输出的目标
	DebugWriter io.Writer = os.Stdout
	
	// 保护并发访问
	mu sync.Mutex
)

// SetDebugMode 设置调试模式
func SetDebugMode(enabled bool) {
	mu.Lock()
	defer mu.Unlock()
	DebugMode = enabled
}

// SetDebugWriter 设置调试输出的目标
func SetDebugWriter(w io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	DebugWriter = w
}

// IncreaseIndent 增加缩进级别
func IncreaseIndent() {
	mu.Lock()
	defer mu.Unlock()
	DebugIndent += 2
}

// DecreaseIndent 减少缩进级别
func DecreaseIndent() {
	mu.Lock()
	defer mu.Unlock()
	if DebugIndent >= 2 {
		DebugIndent -= 2
	}
}

// Print 打印调试信息
func Print(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	
	if !DebugMode {
		return
	}
	
	indent := strings.Repeat(" ", DebugIndent)
	fmt.Fprintf(DebugWriter, indent+format+"\n", args...)
}

// WithIndent 在增加缩进的情况下执行函数
func WithIndent(fn func()) {
	IncreaseIndent()
	defer DecreaseIndent()
	fn()
}

// TraceEnter 跟踪函数进入
func TraceEnter(name string) {
	Print("ENTER: %s", name)
	IncreaseIndent()
}

// TraceExit 跟踪函数退出
func TraceExit(name string) {
	DecreaseIndent()
	Print("EXIT: %s", name)
}

// Trace 跟踪函数执行
func Trace(name string) func() {
	TraceEnter(name)
	return func() {
		TraceExit(name)
	}
}
