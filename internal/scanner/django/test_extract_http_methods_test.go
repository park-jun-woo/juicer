//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestExtractHTTPMethods 테스트
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
