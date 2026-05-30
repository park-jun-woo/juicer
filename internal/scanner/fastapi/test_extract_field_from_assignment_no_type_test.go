//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldFromAssignment_NoType 테스트
package fastapi

import "testing"

func TestExtractFieldFromAssignment_NoType(t *testing.T) {

	src := []byte("class M:\n    count = 5\n")
	root, _ := parsePython(src)
	assigns := findAllByType(root, "assignment")
	f := extractFieldFromAssignment(assigns[0], src)
	if f == nil || f.name != "count" || f.typeName != "" {
		t.Fatalf("got %+v", f)
	}
}
