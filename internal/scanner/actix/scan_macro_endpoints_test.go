//ff:func feature=scan type=test control=sequence topic=actix
//ff:what scanMacroEndpoints — 매크로 라우트에서 엔드포인트 추출을 검증
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestScanMacroEndpoints(t *testing.T) {
	root, err := parseRust([]byte(macroRoutesSource))
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: []byte(macroRoutesSource), root: root}
	files := []*fileInfo{fi}

	sIdx := buildStructIndex(files)
	cache := map[string][]scanner.Field{}
	handlerFuncs := map[string]*handlerInfo{}

	eps := scanMacroEndpoints(files, sIdx, cache, handlerFuncs)
	if len(eps) != 4 {
		t.Fatalf("expected 4 macro endpoints, got %d: %+v", len(eps), eps)
	}
	// collectHandlerFuncs should have populated handler index.
	if len(handlerFuncs) == 0 {
		t.Error("expected handlerFuncs to be populated")
	}
}

func TestScanMacroEndpoints_NoFiles(t *testing.T) {
	eps := scanMacroEndpoints(nil, structIndex{}, map[string][]scanner.Field{}, map[string]*handlerInfo{})
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %+v", eps)
	}
}
