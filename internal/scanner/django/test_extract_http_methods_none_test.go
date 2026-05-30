//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractHTTPMethods_None 테스트
package django

import "testing"

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
