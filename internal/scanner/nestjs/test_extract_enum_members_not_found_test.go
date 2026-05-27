//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what extractEnumMembers 없는 enum 이름 테스트
package nestjs

import "testing"

func TestExtractEnumMembers_NotFound(t *testing.T) {
	src := []byte(`enum Other { A = 'a' }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	vals := extractEnumMembers(root, src, "Missing")
	if vals != nil {
		t.Fatalf("expected nil, got %v", vals)
	}
}
