//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveModelFields_GlobalEntryMissingModel 테스트
package fastapi

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveModelFields_GlobalEntryMissingModel(t *testing.T) {

	fi := &fileInfo{absPath: "/a.py", models: map[string][]pydanticField{}}
	globalModels := map[string]*fileInfo{"Ghost": fi}
	cache := make(map[string][]scanner.Field)
	req := modelRequest{typeName: "Ghost", referrer: "/x/main.py"}
	if got := resolveModelFields(req, cache, globalModels); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
