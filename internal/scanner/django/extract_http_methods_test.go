//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractHTTPMethods — APIView body의 HTTP 메서드 추출 분기를 검증
package django

import "testing"

func TestExtractHTTPMethods(t *testing.T) {
	src := `
class PingView(APIView):
    def get(self, request):
        return Response()
    def post(self, request):
        return Response()
    def helper(self):
        pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if body == nil {
		t.Fatal("no class body")
	}
	methods := extractHTTPMethods(body, []byte(src))
	if len(methods) != 2 {
		t.Fatalf("expected [GET POST], got %v", methods)
	}
	got := map[string]bool{}
	for _, m := range methods {
		got[m] = true
	}
	if !got["GET"] || !got["POST"] {
		t.Errorf("expected GET and POST, got %v", methods)
	}
}

func TestExtractHTTPMethods_None(t *testing.T) {
	src := `
class C:
    def helper(self):
        pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if m := extractHTTPMethods(body, []byte(src)); len(m) != 0 {
		t.Fatalf("expected no methods, got %v", m)
	}
}
