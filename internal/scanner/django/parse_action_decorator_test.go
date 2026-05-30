//ff:func feature=scan type=test control=sequence topic=django
//ff:what parseActionDecorator — @action 데코레이터 파싱 분기를 검증
package django

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func decorators(root *sitter.Node) []*sitter.Node {
	return findAllByType(root, "decorator")
}

func TestParseActionDecorator_Action(t *testing.T) {
	src := `
@action(detail=True, methods=['post'])
def set_password(self):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	decs := decorators(root)
	if len(decs) == 0 {
		t.Fatal("no decorator")
	}
	ai := parseActionDecorator(decs[0], []byte(src))
	if ai == nil {
		t.Fatal("expected actionInfo")
	}
	if !ai.detail {
		t.Error("expected detail true")
	}
	if len(ai.methods) != 1 || ai.methods[0] != "POST" {
		t.Errorf("methods = %v, want [POST]", ai.methods)
	}
}

func TestParseActionDecorator_NotCall(t *testing.T) {
	// @login_required without parentheses has no call child.
	src := `
@login_required
def view(self):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	decs := decorators(root)
	if ai := parseActionDecorator(decs[0], []byte(src)); ai != nil {
		t.Fatalf("expected nil for non-call decorator, got %+v", ai)
	}
}

func TestParseActionDecorator_OtherCall(t *testing.T) {
	// @api_view(...) is a call but not "action".
	src := `
@api_view(['GET'])
def view(self):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	decs := decorators(root)
	if ai := parseActionDecorator(decs[0], []byte(src)); ai != nil {
		t.Fatalf("expected nil for non-action decorator, got %+v", ai)
	}
}
