//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what typeBoxValueToField 스칼라/배열/객체 TypeBox 값 → Field 변환 테스트
package fastify

import "testing"

func TestTypeBoxValueToField(t *testing.T) {
	// scalar
	fi := mustParse(t, []byte(`Type.String()`))
	call := findAllByType(fi.Root, "call_expression")[0]
	f := typeBoxValueToField("name", call, fi.Src)
	if f.Type != "string" || f.Name != "name" {
		t.Errorf("scalar: %+v", f)
	}

	// array of strings
	fi2 := mustParse(t, []byte(`Type.Array(Type.String())`))
	c2 := findAllByType(fi2.Root, "call_expression")[0]
	a := typeBoxValueToField("tags", c2, fi2.Src)
	if a.Type != "string[]" {
		t.Errorf("array: %+v", a)
	}

	// object
	fi3 := mustParse(t, []byte(`Type.Object({ id: Type.Integer() })`))
	c3 := findAllByType(fi3.Root, "call_expression")[0]
	o := typeBoxValueToField("obj", c3, fi3.Src)
	if o.Type != "object" || len(o.Fields) != 1 {
		t.Errorf("object: %+v", o)
	}
}
