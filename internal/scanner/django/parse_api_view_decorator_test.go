//ff:func feature=scan type=test control=sequence topic=django
//ff:what parseAPIViewDecoratorNode — @api_view 데코레이터 파싱 분기를 검증
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

func TestParseAPIViewDecoratorNode_NoListArg(t *testing.T) {
	// @api_view(METHODS) -> args present but no list literal -> default ["GET"].
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
