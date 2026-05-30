//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestBuildMapField_And_ExtractMapFields_Round5 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestBuildMapField_And_ExtractMapFields_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
var M = map[string]any{"name": "x", "age": 5}
`)
	mapLit := firstMapLiteral(info)
	if mapLit == nil {
		t.Fatal("no map literal")
	}
	fields := extractMapFields(mapLit, info)
	if len(fields) != 2 {
		t.Fatalf("fields: %+v", fields)
	}

	if buildMapField(&ast.BasicLit{}, info) != nil {
		t.Fatal("non-KV should be nil")
	}
}
