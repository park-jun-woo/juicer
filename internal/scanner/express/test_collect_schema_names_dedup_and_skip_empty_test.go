//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestCollectSchemaNames_DedupAndSkipEmpty 테스트
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	"testing"
)

func TestCollectSchemaNames_DedupAndSkipEmpty(t *testing.T) {
	routes := []routeInfo{
		{ZodValidators: []zod.ValidatorInfo{
			{SchemaName: "A"},
			{SchemaName: ""},
			{SchemaName: "A"},
			{SchemaName: "B"},
		}},
		{ZodValidators: []zod.ValidatorInfo{
			{SchemaName: "B"},
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
