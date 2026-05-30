//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumAssignmentValue_Irrelevant 테스트
package nestjs

import "testing"

func TestExtractEnumAssignmentValue_Irrelevant(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, _ := parseTypeScript(src)
	nums := findAllByType(root, "number")
	if len(nums) == 0 {
		t.Skip("no number")
	}
	if _, ok := extractEnumAssignmentValue(nums[0], src); ok {
		t.Fatal("expected false for irrelevant node")
	}
}
