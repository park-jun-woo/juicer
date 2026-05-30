//ff:func feature=scan type=test control=sequence topic=hono
//ff:what createRouteObj 테스트 헬퍼
package hono

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// createRouteObj returns the object literal node of the first createRoute({...}) call.
func createRouteObj(t *testing.T, objSrc string) (*fileInfo, *sitter.Node) {
	t.Helper()
	src := []byte("createRoute(" + objSrc + ");\n")
	fi := mustParse(t, src)
	call := findAllByType(fi.Root, "call_expression")[0]
	inner := findChildByType(call, "arguments")
	obj := findChildByType(inner, "object")
	if obj == nil {
		t.Fatal("no object node")
	}
	return fi, obj
}
