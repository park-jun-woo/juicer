//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what TestExtractHTTPMethods req.method === "X" 추출 + OPTIONS 제외 테스트
package supafunc

import (
	"testing"
)

func TestExtractHTTPMethods(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  if (req.method === "OPTIONS") {
    return new Response("ok")
  }
  if (req.method === "GET") {
    return new Response("get")
  }
  if (req.method === "POST") {
    return new Response("post")
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body found")
	}
	methods := extractHTTPMethods(body, fi.Src)
	if len(methods) != 2 {
		t.Fatalf("expected 2 methods, got %d: %v", len(methods), methods)
	}
	found := map[string]bool{}
	for _, m := range methods {
		found[m] = true
	}
	if !found["GET"] {
		t.Fatal("missing GET")
	}
	if !found["POST"] {
		t.Fatal("missing POST")
	}
	if found["OPTIONS"] {
		t.Fatal("OPTIONS should be excluded")
	}
}
