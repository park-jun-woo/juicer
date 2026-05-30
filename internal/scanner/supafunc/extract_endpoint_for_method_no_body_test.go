//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what extractEndpointForMethodNoBody 테스트 (round5)
package supafunc

import "testing"

func TestExtractEndpointForMethodNoBody_Round5(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  if (req.method === "GET") {
    return new Response("ok", { status: 200 })
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}
	blocks := extractMethodBlock(body, fi.Src)
	block := blocks["GET"]
	if block == nil {
		t.Fatal("no GET block")
	}
	ep := extractEndpointForMethodNoBody(block, fi.Src, "GET", "/fn", "handler", "index.ts")
	if ep.Method != "GET" {
		t.Errorf("method: got %s", ep.Method)
	}
	// no-body variant must not set a request body
	if ep.Request != nil && ep.Request.Body != nil {
		t.Errorf("expected no body, got %+v", ep.Request.Body)
	}
}
