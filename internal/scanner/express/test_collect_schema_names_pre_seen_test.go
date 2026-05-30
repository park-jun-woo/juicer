//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectSchemaNames_PreSeen 테스트
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	"testing"
)

func TestCollectSchemaNames_PreSeen(t *testing.T) {
	routes := []routeInfo{{ZodValidators: []zod.ValidatorInfo{{SchemaName: "X"}}}}
	seen := map[string]bool{"X": true}
	var names []string
	collectSchemaNames(routes, seen, &names)
	if len(names) != 0 {
		t.Fatalf("expected none (pre-seen), got %v", names)
	}
}
