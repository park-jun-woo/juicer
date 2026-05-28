//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what extractMethodBlock 단일 메서드 함수에서 빈 맵 반환 테스트
package supafunc

import (
	"testing"
)

func TestExtractMethodBlock_SingleMethod(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const { name } = await req.json()
  return new Response(JSON.stringify({ message: name }), { status: 200 })
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}

	blocks := extractMethodBlock(body, fi.Src)
	if len(blocks) != 0 {
		t.Fatalf("expected 0 blocks for single-method, got %d", len(blocks))
	}
}
