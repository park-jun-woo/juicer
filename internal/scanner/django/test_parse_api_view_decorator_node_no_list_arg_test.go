//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseAPIViewDecoratorNode_NoListArg 테스트
package django

import "testing"

func TestParseAPIViewDecoratorNode_NoListArg(t *testing.T) {

	src := `
@api_view(METHODS)
def view(request):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	decs := decorators(root)
	got := parseAPIViewDecoratorNode(decs[0], []byte(src))
	if len(got) != 1 || got[0] != "GET" {
		t.Fatalf("expected default [GET], got %v", got)
	}
}
