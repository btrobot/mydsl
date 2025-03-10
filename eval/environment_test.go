package eval

import (
	"testing"
)

func TestEnvironment(t *testing.T) {
	env := NewEnvironment()
	
	// 测试设置和获取变量
	env.Set("x", &Integer{Value: 5})
	obj, ok := env.Get("x")
	if !ok {
		t.Fatalf("variable x not found")
	}
	
	integer, ok := obj.(*Integer)
	if !ok {
		t.Fatalf("object is not Integer. got=%T (%+v)", obj, obj)
	}
	
	if integer.Value != 5 {
		t.Errorf("integer.Value wrong. got=%d, want=%d", integer.Value, 5)
	}
	
	// 测试获取不存在的变量
	_, ok = env.Get("y")
	if ok {
		t.Errorf("variable y found, but it shouldn't exist")
	}
}

func TestEnclosedEnvironment(t *testing.T) {
	outer := NewEnvironment()
	outer.Set("x", &Integer{Value: 5})
	
	inner := NewEnclosedEnvironment(outer)
	
	// 测试从外部环境获取变量
	obj, ok := inner.Get("x")
	if !ok {
		t.Fatalf("variable x not found in inner environment")
	}
	
	integer, ok := obj.(*Integer)
	if !ok {
		t.Fatalf("object is not Integer. got=%T (%+v)", obj, obj)
	}
	
	if integer.Value != 5 {
		t.Errorf("integer.Value wrong. got=%d, want=%d", integer.Value, 5)
	}
	
	// 测试在内部环境设置变量
	inner.Set("x", &Integer{Value: 10})
	obj, _ = inner.Get("x")
	integer, _ = obj.(*Integer)
	if integer.Value != 10 {
		t.Errorf("integer.Value wrong. got=%d, want=%d", integer.Value, 10)
	}
	
	// 测试外部环境的变量没有改变
	obj, _ = outer.Get("x")
	integer, _ = obj.(*Integer)
	if integer.Value != 5 {
		t.Errorf("integer.Value in outer env wrong. got=%d, want=%d", integer.Value, 5)
	}
	
	// 测试在内部环境设置新变量
	inner.Set("y", &String{Value: "inner"})
	
	// 测试外部环境不能访问内部变量
	_, ok = outer.Get("y")
	if ok {
		t.Errorf("variable y found in outer environment, but it shouldn't exist there")
	}
}
