//ff:func feature=scan type=extract control=sequence
//ff:what TestApplyFieldTags_JSONWithOmitempty 테스트
package scanner

import "testing"

func TestApplyFieldTags_JSONWithOmitempty(t *testing.T) {
	f := &Field{Name: "Name"}
	excluded := ApplyFieldTags(f, `json:"name,omitempty"`)
	if excluded {
		t.Fatal("should not be excluded")
	}
	if f.JSON != "name" {
		t.Fatalf("expected name, got %s", f.JSON)
	}
}
