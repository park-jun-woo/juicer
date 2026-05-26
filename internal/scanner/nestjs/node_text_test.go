//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what nodeText 테스트
package nestjs

import "testing"

func TestNodeText_Basic(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	text := nodeText(root, src)
	if text != "const x = 1;" {
		t.Fatalf("unexpected: %q", text)
	}
}
