//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseAPIViewDecoratorNode_NotCall 테스트
package django

import "testing"

func TestParseAPIViewDecoratorNode_NotCall(t *testing.T) {
	src := `
@login_required
def view(request):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	decs := decorators(root)
	if got := parseAPIViewDecoratorNode(decs[0], []byte(src)); got != nil {
		t.Fatalf("expected nil for non-call decorator, got %v", got)
	}
}
