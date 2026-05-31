//ff:func feature=scan type=test topic=flask control=sequence
//ff:what assignmentJSONVar JSON 소스 할당의 대상 변수명 추출 테스트
package flask

import "testing"

func TestAssignmentJSONVar(t *testing.T) {
	src := []byte("data = request.get_json()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	asgn := findAllByType(root, "assignment")[0]
	if got := assignmentJSONVar(asgn, src); got != "data" {
		t.Errorf("got %q", got)
	}
	// non-json rhs
	src2 := []byte("data = 5\n")
	root2, _ := parsePython(src2)
	asgn2 := findAllByType(root2, "assignment")[0]
	if got := assignmentJSONVar(asgn2, src2); got != "" {
		t.Errorf("non-json: got %q", got)
	}
}
