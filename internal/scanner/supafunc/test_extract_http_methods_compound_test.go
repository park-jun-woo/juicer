//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what TestExtractHTTPMethodsCompound 복합 binary_expression에서 비메서드 문자열 오인 방지 테스트
package supafunc

import (
	"testing"
)

func TestExtractHTTPMethodsCompound(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  if (req.method === "POST" && functionPath === "/deploy") {
    return new Response("deployed")
  }
  if (req.method === "GET") {
    return new Response("get")
  }
  if (path === "create-scheduling-link") {
    return new Response("ok")
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body found")
	}
	methods := extractHTTPMethods(body, fi.Src)
	found := map[string]bool{}
	for _, m := range methods {
		found[m] = true
	}
	if !found["POST"] {
		t.Fatal("missing POST")
	}
	if !found["GET"] {
		t.Fatal("missing GET")
	}
	if found["/DEPLOY"] {
		t.Fatal("/DEPLOY should not be extracted as a method")
	}
	if found["DEPLOY"] {
		t.Fatal("DEPLOY should not be extracted as a method")
	}
	if found["CREATE-SCHEDULING-LINK"] {
		t.Fatal("CREATE-SCHEDULING-LINK should not be extracted as a method")
	}
	if len(methods) != 2 {
		t.Fatalf("expected 2 methods, got %d: %v", len(methods), methods)
	}
}
