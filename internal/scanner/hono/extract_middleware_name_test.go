//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractMiddlewareName 테스트
package hono

import "testing"

func midArgOf(t *testing.T, src string) (string, string) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	// middleware = second arg (index 1)
	return extractMiddlewareName(nodes[1], fi.Src), ""
}

func TestExtractMiddlewareName_Identifier(t *testing.T) {
	got, _ := midArgOf(t, `app.get("/x", auth, h);`)
	if got != "auth" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractMiddlewareName_Call(t *testing.T) {
	got, _ := midArgOf(t, `app.get("/x", auth(), h);`)
	if got != "auth" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractMiddlewareName_Member(t *testing.T) {
	got, _ := midArgOf(t, `app.get("/x", mw.auth, h);`)
	if got != "mw.auth" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractMiddlewareName_Other(t *testing.T) {
	got, _ := midArgOf(t, `app.get("/x", { a: 1 }, h);`)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
