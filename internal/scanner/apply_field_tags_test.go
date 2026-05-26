//ff:func feature=scan type=test control=sequence
//ff:what TestApplyFieldTags_JSON 테스트
package scanner

import "testing"

func TestApplyFieldTags_JSON(t *testing.T) {
	f := &Field{Name: "Name"}
	excluded := applyFieldTags(f, `json:"name"`)
	if excluded {
		t.Fatal("should not be excluded")
	}
	if f.JSON != "name" {
		t.Fatalf("expected name, got %s", f.JSON)
	}
}
