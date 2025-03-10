package eval

import (
    "bytes"
    "fmt"
    "hash/fnv"
    "strings"
    
    "github.com/btrobot/mydsl/ast"
)

// ObjectType 表示对象类型
type ObjectType string

const (
    NULL_OBJ         = "NULL"
    BOOLEAN_OBJ      = "BOOLEAN"
    INTEGER_OBJ      = "INTEGER"
    FLOAT_OBJ        = "FLOAT"
    STRING_OBJ       = "STRING"
    ARRAY_OBJ        = "ARRAY"
    HASH_OBJ         = "HASH"
    FUNCTION_OBJ     = "FUNCTION"
    BUILTIN_OBJ      = "BUILTIN"
    ERROR_OBJ        = "ERROR"
    RETURN_VALUE_OBJ = "RETURN_VALUE"
    
    // 爬虫相关对象类型
    HTML_DOC_OBJ     = "HTML_DOC"
    SELECTOR_OBJ     = "SELECTOR"
    HTTP_RESPONSE_OBJ = "HTTP_RESPONSE"
)

// Object 表示所有值类型的接口
type Object interface {
    Type() ObjectType
    Inspect() string
}

// Hashable 表示可哈希对象的接口
type Hashable interface {
    HashKey() HashKey
}

// HashKey 表示哈希键
type HashKey struct {
    Type  ObjectType
    Value uint64
}

// Integer 表示整数对象
type Integer struct {
    Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) HashKey() HashKey {
    return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// Float 表示浮点数对象
type Float struct {
    Value float64
}

func (f *Float) Type() ObjectType { return FLOAT_OBJ }
func (f *Float) Inspect() string { return fmt.Sprintf("%g", f.Value) }

// Boolean 表示布尔对象
type Boolean struct {
    Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) HashKey() HashKey {
    var value uint64
    if b.Value {
        value = 1
    } else {
        value = 0
    }
    return HashKey{Type: b.Type(), Value: value}
}

// Null 表示空对象
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }

// String 表示字符串对象
type String struct {
    Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }
func (s *String) HashKey() HashKey {
    h := fnv.New64a()
    h.Write([]byte(s.Value))
    return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// ReturnValue 表示返回值对象
type ReturnValue struct {
    Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error 表示错误对象
type Error struct {
    Message string
    Line    int
    Column  int
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { 
    return fmt.Sprintf("ERROR: %s at line %d, column %d", 
        e.Message, e.Line, e.Column) 
}

// Function 表示函数对象
type Function struct {
    Parameters []*ast.Identifier
    Body       *ast.BlockStatement
    Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
    var out bytes.Buffer
    
    params := []string{}
    for _, p := range f.Parameters {
        params = append(params, p.String())
    }
    
    out.WriteString("function")
    out.WriteString("(")
    out.WriteString(strings.Join(params, ", "))
    out.WriteString(") {\n")
    out.WriteString(f.Body.String())
    out.WriteString("\n}")
    
    return out.String()
}

// BuiltinFunction 表示内置函数类型
type BuiltinFunction func(args ...Object) Object

// Builtin 表示内置函数对象
type Builtin struct {
    Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string { return "builtin function" }

// Array 表示数组对象
type Array struct {
    Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
    var out bytes.Buffer
    
    elements := []string{}
    for _, e := range a.Elements {
        elements = append(elements, e.Inspect())
    }
    
    out.WriteString("[")
    out.WriteString(strings.Join(elements, ", "))
    out.WriteString("]")
    
    return out.String()
}

// HashPair 表示哈希对中的键值对
type HashPair struct {
    Key   Object
    Value Object
}

// Hash 表示哈希对象
type Hash struct {
    Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
    var out bytes.Buffer
    
    pairs := []string{}
    for _, pair := range h.Pairs {
        pairs = append(pairs, fmt.Sprintf("%s: %s",
            pair.Key.Inspect(), pair.Value.Inspect()))
    }
    
    out.WriteString("{")
    out.WriteString(strings.Join(pairs, ", "))
    out.WriteString("}")
    
    return out.String()
}

// HTMLDocument 表示 HTML 文档对象
type HTMLDocument struct {
    Content string
    URL     string
}

func (h *HTMLDocument) Type() ObjectType { return HTML_DOC_OBJ }
func (h *HTMLDocument) Inspect() string { 
    return fmt.Sprintf("HTMLDocument(%s)", h.URL) 
}

// Selector 表示选择器对象
type Selector struct {
    Value string
}

func (s *Selector) Type() ObjectType { return SELECTOR_OBJ }
func (s *Selector) Inspect() string { return fmt.Sprintf("@%s", s.Value) }
func (s *Selector) HashKey() HashKey {
    h := fnv.New64a()
    h.Write([]byte(s.Value))
    return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HTTPResponse 表示 HTTP 响应对象
type HTTPResponse struct {
    StatusCode int
    Body       string
    Headers    map[string]string
    URL        string
}

func (hr *HTTPResponse) Type() ObjectType { return HTTP_RESPONSE_OBJ }
func (hr *HTTPResponse) Inspect() string { 
    return fmt.Sprintf("HTTPResponse(%d, %s)", hr.StatusCode, hr.URL) 
}
