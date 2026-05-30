//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectNeededSchemas_SkipsNoRouters 테스트
package express

import "testing"

func TestCollectNeededSchemas_SkipsNoRouters(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	ctx := &scanContext{
		parsed: map[string]*fileInfo{"a.ts": fi},

		allRouters: map[string]map[string]bool{},
	}
	names := collectNeededSchemas(ctx)
	if len(names) != 0 {
		t.Fatalf("expected no schemas, got %v", names)
	}
}
