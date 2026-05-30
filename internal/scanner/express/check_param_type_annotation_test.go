//ff:func feature=scan type=test control=sequence topic=express
//ff:what checkParamTypeAnnotation — Router 타입 어노테이션 등록 분기를 검증
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstParamOfType(root *sitter.Node, typ string) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == typ {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}

func TestCheckParamTypeAnnotation_Router(t *testing.T) {
	fi := mustParse(t, []byte("function f(r: Router) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	if param == nil {
		t.Fatal("no required_parameter")
	}
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if !routers["r"] {
		t.Fatalf("expected r registered as router, got %v", routers)
	}
}

func TestCheckParamTypeAnnotation_ExpressRouter(t *testing.T) {
	fi := mustParse(t, []byte("function f(r: express.Router) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if !routers["r"] {
		t.Fatalf("expected r registered, got %v", routers)
	}
}

func TestCheckParamTypeAnnotation_OtherType(t *testing.T) {
	fi := mustParse(t, []byte("function f(x: number) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected no routers for non-Router type, got %v", routers)
	}
}

func TestCheckParamTypeAnnotation_NoTypeAnnotation(t *testing.T) {
	// A parameter without a type annotation hits the nil typeAnn early return.
	fi := mustParse(t, []byte("function f(r) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	if param == nil {
		t.Fatal("no required_parameter")
	}
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected no routers for untyped param, got %v", routers)
	}
}

func TestCheckParamTypeAnnotation_NotParameter(t *testing.T) {
	// A non-parameter node returns early.
	fi := mustParse(t, []byte("const x = 1;\n"))
	routers := map[string]bool{}
	checkParamTypeAnnotation(fi.Root, fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected no routers, got %v", routers)
	}
}
