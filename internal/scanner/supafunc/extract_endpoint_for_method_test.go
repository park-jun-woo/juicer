//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what extractEndpointForMethod 테스트 (round5)
package supafunc

import "testing"

func TestExtractEndpointForMethod_Round5(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  if (req.method === "POST") {
    const { name } = await req.json()
    return new Response(JSON.stringify({ ok: true }), { status: 201 })
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}
	blocks := extractMethodBlock(body, fi.Src)
	block := blocks["POST"]
	if block == nil {
		t.Fatal("no POST block")
	}
	ep := extractEndpointForMethod(block, fi.Src, "POST", "/fn", "handler", "index.ts")
	if ep.Method != "POST" {
		t.Errorf("method: got %s", ep.Method)
	}
	if ep.Path != "/fn" {
		t.Errorf("path: got %s", ep.Path)
	}
	if ep.Handler != "handler" {
		t.Errorf("handler: got %s", ep.Handler)
	}
	if ep.File != "index.ts" {
		t.Errorf("file: got %s", ep.File)
	}
}
