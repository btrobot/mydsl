package eval

// Environment 表示执行环境
type Environment struct {
    store map[string]Object
    outer *Environment
}

// NewEnvironment 创建新的环境
func NewEnvironment() *Environment {
    s := make(map[string]Object)
    return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment 创建嵌套环境
func NewEnclosedEnvironment(outer *Environment) *Environment {
    env := NewEnvironment()
    env.outer = outer
    return env
}

// Get 获取变量值
func (e *Environment) Get(name string) (Object, bool) {
    obj, ok := e.store[name]
    if !ok && e.outer != nil {
        obj, ok = e.outer.Get(name)
    }
    return obj, ok
}

// Set 设置变量值
func (e *Environment) Set(name string, val Object) Object {
    e.store[name] = val
    return val
}

// GetAll 获取所有变量
func (e *Environment) GetAll() map[string]Object {
    return e.store
}
