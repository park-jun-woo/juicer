//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveModelFields 테스트
package fastapi

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveModelFields(t *testing.T) {
	fields := []pydanticField{{name: "email", typeName: "str"}}
	fi := &fileInfo{absPath: "/a.py", models: map[string][]pydanticField{"User": fields}}
	globalModels := map[string]*fileInfo{"User": fi}
	cache := make(map[string][]scanner.Field)

	req := modelRequest{typeName: "User"}
	result := resolveModelFields(req, cache, globalModels)
	if len(result) != 1 {
		t.Fatalf("expected 1 field, got %d", len(result))
	}

	result2 := resolveModelFields(req, cache, globalModels)
	if len(result2) != 1 {
		t.Fatalf("cache: expected 1 field, got %d", len(result2))
	}

	req2 := modelRequest{typeName: "Unknown"}
	result3 := resolveModelFields(req2, cache, globalModels)
	if result3 != nil {
		t.Fatalf("expected nil for unknown, got %v", result3)
	}
}
