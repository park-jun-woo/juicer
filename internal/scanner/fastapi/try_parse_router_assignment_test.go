//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what tryParseRouterAssignment 테스트
package fastapi

import "testing"

func TestTryParseRouterAssignment(t *testing.T) {
	src := []byte("app = FastAPI()\nrouter = APIRouter(prefix='/api')\nother = SomeClass()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	assigns := findAllByType(root, "assignment")
	found := 0
	for _, a := range assigns {
		ri := tryParseRouterAssignment(a, src)
		if ri != nil {
			found++
		}
	}
	if found < 2 {
		t.Fatalf("expected at least 2 router assignments, got %d", found)
	}
}

func TestTryParseRouterAssignment_NonRouter(t *testing.T) {
	// assignment with no call on the right -> nil (call == nil)
	src := []byte("x = 5\n")
	root, _ := parsePython(src)
	for _, a := range findAllByType(root, "assignment") {
		if ri := tryParseRouterAssignment(a, src); ri != nil {
			t.Fatalf("non-call assignment should be nil, got %v", ri)
		}
	}
}

func TestTryParseRouterAssignment_NotRouterClass(t *testing.T) {
	// call but func name is not a router class -> nil
	src := []byte("db = SomeClass()\n")
	root, _ := parsePython(src)
	for _, a := range findAllByType(root, "assignment") {
		if ri := tryParseRouterAssignment(a, src); ri != nil {
			t.Fatalf("non-router class should be nil, got %v", ri)
		}
	}
}

func TestTryParseRouterAssignment_NoArgs(t *testing.T) {
	// FastAPI() with no prefix arg still produces a routerInfo
	src := []byte("app = FastAPI()\n")
	root, _ := parsePython(src)
	a := findAllByType(root, "assignment")[0]
	ri := tryParseRouterAssignment(a, src)
	if ri == nil || !ri.isFastAPI || ri.varName != "app" {
		t.Fatalf("FastAPI assignment: got %v", ri)
	}
	if ri.prefix != "" {
		t.Fatalf("expected empty prefix, got %q", ri.prefix)
	}
}
