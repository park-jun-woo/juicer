//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseAPIViewDecoratorNode_OtherCall 테스트
package django

import "testing"

func TestParseAPIViewDecoratorNode_OtherCall(t *testing.T) {
	src := `
@action(detail=True)
def view(self):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	decs := decorators(root)
	if got := parseAPIViewDecoratorNode(decs[0], []byte(src)); got != nil {
		t.Fatalf("expected nil for non-api_view decorator, got %v", got)
	}
}
