//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseActionDecorator_Action 테스트
package django

import "testing"

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
