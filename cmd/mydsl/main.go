package main

import (
	"flag"
	"fmt"
	"os"
	
	"github.com/btrobot/mydsl/internal/debug"
)

var (
	debugMode = flag.Bool("debug", false, "Enable debug mode")
	version   = flag.Bool("version", false, "Show version information")
)

const (
	VERSION = "0.1.0"
)

func main() {
	flag.Parse()
	
	if *version {
		fmt.Printf("MyDSL version %s\n", VERSION)
		return
	}
	
	// 设置调试模式
	debug.SetDebugMode(*debugMode)
	
	// 获取输入文件
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: mydsl [options] <filename>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	filename := args[0]
	
	// 读取文件内容
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	// 这里将来会添加词法分析、语法分析和解释执行的代码
	fmt.Printf("Read %d bytes from %s\n", len(content), filename)
	
	if *debugMode {
		debug.Print("Debug mode enabled")
		debug.Print("File content: %s", string(content))
	}
}
