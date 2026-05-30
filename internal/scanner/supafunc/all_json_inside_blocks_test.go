//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what allJSONInsideBlocks 테스트 (round5)
package supafunc

import "testing"

func TestAllJSONInsideBlocks_NoJSON_Round5(t *testing.T) {
	// No req.json() anywhere => returns true (no body outside blocks).
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
	if !allJSONInsideBlocks(body, fi.Src, blocks) {
		t.Fatal("expected true when no req.json()")
	}
}

func TestAllJSONInsideBlocks_Inside_Round5(t *testing.T) {
	// req.json() inside a method block => true.
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
	if !allJSONInsideBlocks(body, fi.Src, blocks) {
		t.Fatal("expected true when req.json() inside block")
	}
}

func TestAllJSONInsideBlocks_Outside_Round5(t *testing.T) {
	// req.json() at top-level (outside any method block) => false.
	src := []byte(`
serve(async (req) => {
  const { name } = await req.json()
  if (req.method === "POST") {
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
	if allJSONInsideBlocks(body, fi.Src, blocks) {
		t.Fatal("expected false when req.json() outside blocks")
	}
}
