//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestMergeFields 테스트
package nestjs

import "testing"

func TestMergeFields(t *testing.T) {
	parent := []dtoField{
		{name: "name", tsType: "string"},
		{name: "email", tsType: "string"},
	}
	child := []dtoField{
		{name: "email", tsType: "string", optional: true},
		{name: "id", tsType: "number"},
	}
	result := mergeFields(parent, child)
	if len(result) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(result))
	}
	names := make(map[string]bool)
	for _, f := range result {
		names[f.name] = true
	}
	for _, want := range []string{"name", "email", "id"} {
		if !names[want] {
			t.Fatalf("missing field %q", want)
		}
	}
}
