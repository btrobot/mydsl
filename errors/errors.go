package errors

import "fmt"

// ErrorType 表示错误类型
type ErrorType int

const (
    SyntaxError ErrorType = iota
    RuntimeError
    NetworkError
    SelectorError
    TypeError
    ReferenceError
)

// Error 表示 DSL 错误
type Error struct {
    Type    ErrorType // 错误类型
    Message string    // 错误消息
    Line    int       // 行号
    Column  int       // 列号
    Value   interface{} // 相关值（可选）
}

// Error 实现 error 接口
func (e *Error) Error() string {
    return fmt.Sprintf("%s at line %d, column %d: %s", 
        e.TypeString(), e.Line, e.Column, e.Message)
}

// TypeString 返回错误类型的字符串表示
func (e *Error) TypeString() string {
    switch e.Type {
    case SyntaxError:
        return "Syntax Error"
    case RuntimeError:
        return "Runtime Error"
    case NetworkError:
        return "Network Error"
    case SelectorError:
        return "Selector Error"
    case TypeError:
        return "Type Error"
    case ReferenceError:
        return "Reference Error"
    default:
        return "Unknown Error"
    }
}

// NewSyntaxError 创建语法错误
func NewSyntaxError(message string, line, column int) *Error {
    return &Error{
        Type:    SyntaxError,
        Message: message,
        Line:    line,
        Column:  column,
    }
}

// NewRuntimeError 创建运行时错误
func NewRuntimeError(message string, line, column int, value interface{}) *Error {
    return &Error{
        Type:    RuntimeError,
        Message: message,
        Line:    line,
        Column:  column,
        Value:   value,
    }
}

// NewNetworkError 创建网络错误
func NewNetworkError(message string, line, column int) *Error {
    return &Error{
        Type:    NetworkError,
        Message: message,
        Line:    line,
        Column:  column,
    }
}

// NewSelectorError 创建选择器错误
func NewSelectorError(message string, line, column int) *Error {
    return &Error{
        Type:    SelectorError,
        Message: message,
        Line:    line,
        Column:  column,
    }
}

// NewTypeError 创建类型错误
func NewTypeError(message string, line, column int, value interface{}) *Error {
    return &Error{
        Type:    TypeError,
        Message: message,
        Line:    line,
        Column:  column,
        Value:   value,
    }
}

// NewReferenceError 创建引用错误
func NewReferenceError(message string, line, column int, value interface{}) *Error {
    return &Error{
        Type:    ReferenceError,
        Message: message,
        Line:    line,
        Column:  column,
        Value:   value,
    }
}
