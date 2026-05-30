//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractActions_None 테스트
package django

import "testing"

func TestExtractActions_None(t *testing.T) {
	src := `
class C:
    def plain(self):
        pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if a := extractActions(body, []byte(src), "v.py"); len(a) != 0 {
		t.Fatalf("expected no actions, got %d", len(a))
	}
}
