//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectRouterParamsFromFunc: formal_parameters 유무 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstFuncDecl(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	fns := findAllByType(fi.Root, "function_declaration")
	if len(fns) == 0 {
		t.Fatal("no function_declaration")
	}
	return fns[0]
}

func TestCollectRouterParamsFromFunc_WithRouterParam(t *testing.T) {
	fi := mustParse(t, []byte(`function setup(r: Router) {}`))
	routers := map[string]bool{}
	collectRouterParamsFromFunc(firstFuncDecl(t, fi), fi.Src, routers)
	if !routers["r"] {
		t.Fatalf("expected r, got %v", routers)
	}
}

func TestCollectRouterParamsFromFunc_NoRouterParam(t *testing.T) {
	fi := mustParse(t, []byte(`function setup(n: number) {}`))
	routers := map[string]bool{}
	collectRouterParamsFromFunc(firstFuncDecl(t, fi), fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}

func TestCollectRouterParamsFromFunc_NoParams(t *testing.T) {
	// an arrow function without formal_parameters node (single ident param)
	fi := mustParse(t, []byte(`const f = x => x;`))
	arrows := findAllByType(fi.Root, "arrow_function")
	if len(arrows) == 0 {
		t.Fatal("no arrow_function")
	}
	routers := map[string]bool{}
	collectRouterParamsFromFunc(arrows[0], fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}
