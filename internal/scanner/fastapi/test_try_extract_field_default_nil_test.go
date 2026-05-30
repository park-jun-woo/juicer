//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestTryExtractField_DefaultNil 테스트
package fastapi

import "testing"

func TestTryExtractField_DefaultNil(t *testing.T) {

	src := []byte("class M(BaseModel):\n    pass\n")
	root, _ := parsePython(src)
	for _, ps := range findAllByType(root, "pass_statement") {
		if f := tryExtractField(ps, src); f != nil {
			t.Fatalf("pass should yield nil, got %v", f)
		}
		return
	}
	t.Skip("no pass_statement")
}
