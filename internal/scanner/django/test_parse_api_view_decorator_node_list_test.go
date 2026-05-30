//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseAPIViewDecoratorNode_List 테스트
package django

import "testing"

func TestParseAPIViewDecoratorNode_List(t *testing.T) {
	src := `
@api_view(['GET', 'POST'])
def view(request):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	decs := decorators(root)
	got := parseAPIViewDecoratorNode(decs[0], []byte(src))
	if len(got) != 2 {
		t.Fatalf("expected 2 methods, got %v", got)
	}
}
