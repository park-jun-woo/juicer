//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestScanBuilderEndpoints 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
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
