//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what extractMethodBlockFromIf 테스트 (round5)
package supafunc

import "testing"

func TestExtractMethodBlockFromIf_Round5(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  if (req.method === "POST") {
    const { name } = await req.json()
    return new Response("ok", { status: 201 })
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}
	blocks := extractMethodBlock(body, fi.Src)
	if blocks["POST"] == nil {
		t.Fatalf("expected POST block, got %v keys", blocks)
	}
}

func TestExtractMethodBlockFromIf_NoMethodCond_Round5(t *testing.T) {
	// if condition unrelated to method => no block registered
	src := []byte(`
serve(async (req) => {
  if (req.url === "/x") {
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
	if len(blocks) != 0 {
		t.Fatalf("expected no blocks for non-method if, got %v", blocks)
	}
}
