package eval

import (
	"strings"
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}
	
	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}
	
	if true1.HashKey() != true2.HashKey() {
		t.Errorf("trues have different hash keys")
	}
	
	if false1.HashKey() != false2.HashKey() {
		t.Errorf("falses have different hash keys")
	}
	
	if true1.HashKey() == false1.HashKey() {
		t.Errorf("true and false have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}
	two1 := &Integer{Value: 2}
	two2 := &Integer{Value: 2}
	
	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}
	
	if two1.HashKey() != two2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}
	
	if one1.HashKey() == two1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}

func TestInspectMethods(t *testing.T) {
	tests := []struct {
		obj      Object
		expected string
	}{
		{&Integer{Value: 5}, "5"},
		{&Float{Value: 3.14}, "3.14"},
		{&Boolean{Value: true}, "true"},
		{&String{Value: "Hello"}, "Hello"},
		{&Null{}, "null"},
		{&Array{Elements: []Object{
			&Integer{Value: 1},
			&Integer{Value: 2},
		}}, "[1, 2]"},
		{&Error{Message: "error", Line: 1, Column: 2}, "ERROR: error at line 1, column 2"},
	}
	
	for _, tt := range tests {
		if tt.obj.Inspect() != tt.expected {
			t.Errorf("obj.Inspect() wrong. got=%q, want=%q", 
				tt.obj.Inspect(), tt.expected)
		}
	}
}

func TestHashInspect(t *testing.T) {
	hash := &Hash{
		Pairs: map[HashKey]HashPair{
			(&String{Value: "name"}).HashKey(): {
				Key:   &String{Value: "name"},
				Value: &String{Value: "John"},
			},
			(&String{Value: "age"}).HashKey(): {
				Key:   &String{Value: "age"},
				Value: &Integer{Value: 30},
			},
		},
	}
	
	// 由于哈希表的顺序不确定，我们只检查长度
	inspect := hash.Inspect()
	if len(inspect) < 10 {
		t.Errorf("hash.Inspect() too short. got=%q", inspect)
	}
	
	// 检查是否包含所有键值对
	if !strings.Contains(inspect, "name") || !strings.Contains(inspect, "John") {
		t.Errorf("hash.Inspect() missing key-value pair. got=%q", inspect)
	}
	
	if !strings.Contains(inspect, "age") || !strings.Contains(inspect, "30") {
		t.Errorf("hash.Inspect() missing key-value pair. got=%q", inspect)
	}
}
