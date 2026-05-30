//ff:func feature=scan type=test control=sequence
//ff:what TestResolveBindType_NoArgsOrNilInfo 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveBindType_NoArgsOrNilInfo(t *testing.T) {
	if tn, f := resolveBindType(&ast.CallExpr{}, newEmptyInfo()); tn != "" || f != nil {
		t.Fatalf("no args: got %q %v", tn, f)
	}
	call := parseCall(t, "c.BodyParser(&req)")
	if tn, f := resolveBindType(call, nil); tn != "" || f != nil {
		t.Fatalf("nil info: got %q %v", tn, f)
	}
}
