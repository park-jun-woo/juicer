//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what typeBoxPairToField pair → Field(required/Optional) 변환 테스트
package fastify

import "testing"

func TestTypeBoxPairToField(t *testing.T) {
	// required scalar
	fi := mustParse(t, []byte(`const s = { name: Type.String() };`))
	pair := findAllByType(fi.Root, "pair")[0]
	f := typeBoxPairToField(pair, fi.Src)
	if f == nil || f.Name != "name" || f.Validate != "required" {
		t.Fatalf("required: %+v", f)
	}

	// optional field
	fi2 := mustParse(t, []byte(`const s = { bio: Type.Optional(Type.String()) };`))
	p2 := findAllByType(fi2.Root, "pair")[0]
	f2 := typeBoxPairToField(p2, fi2.Src)
	if f2 == nil || f2.Validate == "required" {
		t.Errorf("optional should not be required: %+v", f2)
	}
}
