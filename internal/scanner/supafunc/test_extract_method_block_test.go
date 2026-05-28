//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what extractMethodBlock if 분기에서 메서드별 블록 추출 테스트
package supafunc

import (
	"testing"
)

func TestExtractMethodBlock_IfBranches(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  if (req.method === "GET") {
    return new Response("get", { status: 200 })
  }
  if (req.method === "POST") {
    const { name } = await req.json()
    return new Response("post", { status: 201 })
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}

	blocks := extractMethodBlock(body, fi.Src)
	if len(blocks) != 2 {
		t.Fatalf("expected 2 blocks, got %d", len(blocks))
	}
	if blocks["GET"] == nil {
		t.Fatal("missing GET block")
	}
	if blocks["POST"] == nil {
		t.Fatal("missing POST block")
	}
	if blocks["OPTIONS"] != nil {
		t.Fatal("OPTIONS should not be included")
	}
}
