//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what parseCSharp C# 소스 파싱 후 루트 노드 반환 테스트
package dotnet

import "testing"

func TestParseCSharp(t *testing.T) {
	root, err := parseCSharp([]byte(`class C { void M() {} }`))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if root == nil {
		t.Fatal("nil root")
	}
	if root.Type() != "compilation_unit" {
		t.Errorf("root type = %q", root.Type())
	}
}
