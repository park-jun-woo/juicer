//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractPydanticFields_NoBlock 테스트
package fastapi

import "testing"

func TestExtractPydanticFields_NoBlock(t *testing.T) {

	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	if fields := extractPydanticFields(root, src); fields != nil {
		t.Fatalf("expected nil, got %+v", fields)
	}
}
