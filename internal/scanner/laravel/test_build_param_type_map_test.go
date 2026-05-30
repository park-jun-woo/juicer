//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildParamTypeMap 테스트
package laravel

import "testing"

func TestBuildParamTypeMap(t *testing.T) {
	cm := &controllerMethod{params: []methodParam{
		{name: "id", typeName: "int"},
		{name: "name", typeName: "string"},
		{name: "request", typeName: "Request"},
		{name: "untyped", typeName: ""},
		{name: "weird", typeName: "SomeClass"},
	}}
	m := buildParamTypeMap(cm)
	if m["id"] != "integer" || m["name"] != "string" {
		t.Fatalf("got %v", m)
	}
	if _, ok := m["request"]; ok {
		t.Fatal("request should be skipped")
	}
	if _, ok := m["untyped"]; ok {
		t.Fatal("untyped should be skipped")
	}
	if _, ok := m["weird"]; ok {
		t.Fatal("unmapped type should be skipped")
	}
}
