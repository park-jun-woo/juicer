//ff:func feature=scan type=test control=sequence
//ff:what TestResolveBindType_EmptyArgs 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveBindType_EmptyArgs(t *testing.T) {
	call := &ast.CallExpr{}
	info := &types.Info{}
	name, fields := resolveBindType(call, info)
	if name != "" || fields != nil {
		t.Fatal("expected empty")
	}
}

