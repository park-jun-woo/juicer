//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryExtractField_Assignment 테스트
package fastapi

import "testing"

func TestTryExtractField_Assignment(t *testing.T) {

	src := []byte("class M(BaseModel):\n    age: int = 25\n")
	root, _ := parsePython(src)
	as := findAllByType(root, "assignment")
	if len(as) == 0 {
		t.Skip("no assignment node")
	}
	f := tryExtractField(as[0], src)
	if f == nil || f.name != "age" {
		t.Fatalf("assignment field: got %v", f)
	}
}
