//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractOneRoute 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func routeCalls(t *testing.T, src string) (*fileInfo, []*sitter.Node) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	return fi, findAllByType(fi.Root, "call_expression")
}

func TestExtractOneRoute_Success(t *testing.T) {
	fi, calls := routeCalls(t, `app.get("/users/:id", handler);`+"\n")
	inst := map[string]bool{"app": true}
	var got *routeInfo
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			got = ri
		}
	}
	if got == nil {
		t.Fatal("expected route")
	}
	if got.Method != "GET" || got.Path != "/users/:id" || got.Handler != "handler" {
		t.Fatalf("route = %+v", got)
	}
}

func TestExtractOneRoute_NotInstance(t *testing.T) {
	fi, calls := routeCalls(t, `other.get("/x", h);`+"\n")
	inst := map[string]bool{"app": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("non-instance should yield nil, got %+v", ri)
		}
	}
}

func TestExtractOneRoute_NotHTTPMethod(t *testing.T) {
	fi, calls := routeCalls(t, `app.listen("/x", h);`+"\n")
	inst := map[string]bool{"app": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("listen should yield nil, got %+v", ri)
		}
	}
}

func TestExtractOneRoute_NoMemberExpression(t *testing.T) {
	fi, calls := routeCalls(t, `get("/x", h);`+"\n")
	inst := map[string]bool{"app": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("plain call should yield nil, got %+v", ri)
		}
	}
}

func TestExtractOneRoute_TooFewArgs(t *testing.T) {
	fi, calls := routeCalls(t, `app.get("/x");`+"\n")
	inst := map[string]bool{"app": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("single arg should yield nil, got %+v", ri)
		}
	}
}

func TestExtractOneRoute_TemplatePath(t *testing.T) {
	// template_string path is accepted
	fi, calls := routeCalls(t, "app.post(`/users`, h);\n")
	inst := map[string]bool{"app": true}
	var got *routeInfo
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			got = ri
		}
	}
	if got == nil || got.Method != "POST" {
		t.Fatalf("template path route = %+v", got)
	}
}

func TestExtractOneRoute_PathNotString(t *testing.T) {
	fi, calls := routeCalls(t, `app.get(pathVar, h);`+"\n")
	inst := map[string]bool{"app": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("non-string path should yield nil, got %+v", ri)
		}
	}
}
