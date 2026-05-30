//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseMacroAttribute_And_BuildMacroEndpoint_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestParseMacroAttribute_And_BuildMacroEndpoint_Round5(t *testing.T) {
	fi := aFi(t, macroSrc)
	attr := aFirst(t, fi.root, "attribute_item")
	mr := parseMacroAttribute(attr, fi.src)
	if mr == nil {
		t.Fatal("expected macro route")
	}
	if mr.method != "GET" || mr.path != "/health" {
		t.Fatalf("macro route: %+v", mr)
	}
	mr.file = fi
	mr.funcNode = aFirst(t, fi.root, "function_item")
	mr.handler = "health"
	ep := buildMacroEndpoint(*mr, fi, structIndex{}, map[string][]scanner.Field{})
	if ep.Method != "GET" || ep.Path != "/health" {
		t.Fatalf("endpoint: %+v", ep)
	}
}
