//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what typeBoxPairToParam pair → Param(타입/default/Optional) 변환 테스트
package fastify

import "testing"

func TestTypeBoxPairToParam(t *testing.T) {
	fi := mustParse(t, []byte(`const q = { page: Type.Integer({ default: '1' }) };`))
	pair := findAllByType(fi.Root, "pair")[0]
	p := typeBoxPairToParam(pair, fi.Src)
	if p == nil || p.Name != "page" || p.Type != "integer" || p.Default != "1" {
		t.Fatalf("got %+v", p)
	}

	// optional
	fi2 := mustParse(t, []byte(`const q = { sort: Type.Optional(Type.String()) };`))
	p2 := typeBoxPairToParam(findAllByType(fi2.Root, "pair")[0], fi2.Src)
	if p2 == nil || p2.Name != "sort" || p2.Type != "string" {
		t.Errorf("optional: %+v", p2)
	}
}
