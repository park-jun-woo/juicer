//ff:func feature=scan type=test control=sequence topic=actix
//ff:what scanBuilderEndpoints — 빌더 라우트에서 엔드포인트 추출을 검증
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestScanBuilderEndpoints(t *testing.T) {
	root, err := parseRust([]byte(builderRoutesSource))
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: []byte(builderRoutesSource), root: root}
	files := []*fileInfo{fi}

	sIdx := buildStructIndex(files)
	cache := map[string][]scanner.Field{}
	handlerFuncs := map[string]*handlerInfo{}

	eps := scanBuilderEndpoints(files, sIdx, cache, handlerFuncs)
	if len(eps) == 0 {
		t.Fatalf("expected builder endpoints, got none")
	}
	// At least one /api/users route should be present.
	found := false
	for _, ep := range eps {
		if ep.Path == "/api/users" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected /api/users endpoint, got %+v", eps)
	}
}

func TestScanBuilderEndpoints_NoFiles(t *testing.T) {
	eps := scanBuilderEndpoints(nil, structIndex{}, map[string][]scanner.Field{}, map[string]*handlerInfo{})
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %+v", eps)
	}
}
