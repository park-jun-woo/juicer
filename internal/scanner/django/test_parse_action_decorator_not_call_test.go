//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseActionDecorator_NotCall 테스트
package django

import "testing"

func TestParseActionDecorator_NotCall(t *testing.T) {

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
