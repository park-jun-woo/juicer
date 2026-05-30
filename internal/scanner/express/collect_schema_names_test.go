//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectSchemaNames: 빈 이름/중복/신규 수집 분기 검증
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

func TestCollectSchemaNames_DedupAndSkipEmpty(t *testing.T) {
	routes := []routeInfo{
		{ZodValidators: []zod.ValidatorInfo{
			{SchemaName: "A"},
			{SchemaName: ""},  // skipped: empty
			{SchemaName: "A"}, // skipped: duplicate
			{SchemaName: "B"},
		}},
		{ZodValidators: []zod.ValidatorInfo{
			{SchemaName: "B"}, // duplicate across routes
			{SchemaName: "C"},
		}},
	}
	seen := map[string]bool{}
	var names []string
	collectSchemaNames(routes, seen, &names)
	want := []string{"A", "B", "C"}
	if len(names) != len(want) {
		t.Fatalf("got %v, want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Fatalf("got %v, want %v", names, want)
		}
	}
}

func TestCollectSchemaNames_PreSeen(t *testing.T) {
	routes := []routeInfo{{ZodValidators: []zod.ValidatorInfo{{SchemaName: "X"}}}}
	seen := map[string]bool{"X": true}
	var names []string
	collectSchemaNames(routes, seen, &names)
	if len(names) != 0 {
		t.Fatalf("expected none (pre-seen), got %v", names)
	}
}
