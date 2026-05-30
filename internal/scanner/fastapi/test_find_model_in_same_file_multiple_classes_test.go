//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindModelInSameFile_MultipleClasses 테스트
package fastapi

import "testing"

func TestFindModelInSameFile_MultipleClasses(t *testing.T) {
	src := []byte("class A(BaseModel):\n    x: int\nclass B(BaseModel):\n    y: int\n")
	root, _ := parsePython(src)
	if !findModelInSameFile(root, src, "B") {
		t.Fatal("expected B (skips A first)")
	}
}
