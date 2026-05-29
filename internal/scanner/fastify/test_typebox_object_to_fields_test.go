//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TypeBox Type.Array/Type.Object 중첩이 Field로 변환되는지 검증한다
package fastify

import "testing"

func TestTypeBoxObjectToFields(t *testing.T) {
	src := []byte(`const S = Type.Object({
  tags: Type.Array(Type.String()),
  meta: Type.Object({ key: Type.String() }),
  count: Type.Integer({ minimum: 1 })
});`)
	fi := mustParse(t, src)
	vars := extractTypeBoxVars(fi)
	obj := vars["S"]
	if obj == nil {
		t.Fatal("expected S in typebox vars")
	}
	fields := typeBoxObjectToFields(obj, fi.Src)
	fm := make(map[string]int)
	for i, f := range fields {
		fm[f.Name] = i
	}
	if i, ok := fm["tags"]; !ok || fields[i].Type != "string[]" {
		t.Errorf("tags: want string[], got %v", fields)
	}
	if i, ok := fm["meta"]; !ok {
		t.Error("missing meta")
	} else {
		if fields[i].Type != "object" || len(fields[i].Fields) != 1 {
			t.Errorf("meta: want object with 1 nested field, got type=%s fields=%d", fields[i].Type, len(fields[i].Fields))
		}
	}
	if i, ok := fm["count"]; !ok || fields[i].Type != "integer" {
		t.Errorf("count: want integer")
	} else if fields[i].Minimum == nil || *fields[i].Minimum != 1 {
		t.Errorf("count.Minimum: want 1, got %v", fields[i].Minimum)
	}
}
