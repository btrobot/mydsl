package token

import "fmt"

// TokenType 是词法单元类型的字符串表示
type TokenType string

// Token 表示词法单元
type Token struct {
    Type    TokenType // 词法单元类型
    Literal string    // 词法单元的字面值
    Line    int       // 行号
    Column  int       // 列号
}

// String 返回 Token 的字符串表示
func (t Token) String() string {
    return fmt.Sprintf("Token(%s, %q, line:%d, col:%d)", 
        t.Type, t.Literal, t.Line, t.Column)
}

// 定义所有词法单元类型常量
const (
    ILLEGAL = "ILLEGAL" // 非法字符
    EOF     = "EOF"     // 文件结束

    // 标识符和字面量
    IDENT  = "IDENT"  // 标识符
    SYSVAR = "SYSVAR" // 系统变量 ($开头)
    INT    = "INT"    // 整数
    FLOAT  = "FLOAT"  // 浮点数
    STRING = "STRING" // 字符串
    
    // 运算符
    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"
    PERCENT  = "%"
    
    // 比较运算符
    EQ     = "=="
    NOT_EQ = "!="
    LT     = "<"
    GT     = ">"
    LTE    = "<="
    GTE    = ">="
    
    // 逻辑运算符
    AND = "&&"
    OR  = "||"
    
    // 分隔符
    COMMA     = ","
    SEMICOLON = ";"
    COLON     = ":"
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"
    LBRACKET  = "["
    RBRACKET  = "]"
    DOT       = "."
    
    // 特殊运算符
    QUESTION      = "?"
    NULL_COALESCE = "??"
    OPTIONAL_DOT  = "?."
    PIPE          = "|"
    AT            = "@"
    ARROW         = "=>"
    
    // 关键字
    FUNCTION = "FUNCTION"
    LET      = "LET"
    CONST    = "CONST"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"
    WHILE    = "WHILE"
    FOR      = "FOR"
    IN       = "IN"
    BREAK    = "BREAK"
    CONTINUE = "CONTINUE"
    NULL     = "NULL"
    
    // 爬虫相关关键字
    OPEN    = "OPEN"
    EXTRACT = "EXTRACT"
    COLLECT = "COLLECT"
    KEYS    = "KEYS"
    VALUES  = "VALUES"
    LENGTH  = "LENGTH"
    DELETE  = "DELETE"
    
    // 日志相关关键字
    LOG   = "LOG"
    DEBUG = "DEBUG"
    INFO  = "INFO"
    WARN  = "WARN"
    ERROR = "ERROR"
)

// 关键字映射表
var keywords = map[string]TokenType{
    "function": FUNCTION,
    "let":      LET,
    "const":    CONST,
    "true":     TRUE,
    "false":    FALSE,
    "if":       IF,
    "else":     ELSE,
    "return":   RETURN,
    "while":    WHILE,
    "for":      FOR,
    "in":       IN,
    "break":    BREAK,
    "continue": CONTINUE,
    "null":     NULL,
    "open":     OPEN,
    "extract":  EXTRACT,
    "collect":  COLLECT,
    "keys":     KEYS,
    "values":   VALUES,
    "length":   LENGTH,
    "delete":   DELETE,
    "log":      LOG,
    "debug":    DEBUG,
    "info":     INFO,
    "warn":     WARN,
    "error":    ERROR,
}

// LookupIdent 检查标识符是否为关键字
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

// NewToken 创建一个新的词法单元
func NewToken(tokenType TokenType, literal string, line, column int) Token {
    return Token{
        Type:    tokenType,
        Literal: literal,
        Line:    line,
        Column:  column,
    }
}
