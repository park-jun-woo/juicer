//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what mergeParentAndOwnFields: 자식 override / 비중복 부모 보존
package fastapi

import "testing"

func TestMergeParentAndOwnFields(t *testing.T) {
	parent := []pydanticField{
		{name: "id", typeName: "int"},
		{name: "name", typeName: "str"},
	}
	own := []pydanticField{
		{name: "name", typeName: "int"}, // overrides parent's name
		{name: "extra", typeName: "str"},
	}
	merged := mergeParentAndOwnFields(parent, own)
	// id (parent), name (own, overriding), extra (own)
	if len(merged) != 3 {
		t.Fatalf("expected 3, got %d: %+v", len(merged), merged)
	}
	byName := map[string]string{}
	for _, f := range merged {
		byName[f.name] = f.typeName
	}
	if byName["id"] != "int" {
		t.Errorf("id type: %s", byName["id"])
	}
	if byName["name"] != "int" {
		t.Errorf("name should be overridden to int, got %s", byName["name"])
	}
	if byName["extra"] != "str" {
		t.Errorf("extra missing: %s", byName["extra"])
	}
}
