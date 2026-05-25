//ff:func feature=scan type=extract control=sequence
//ff:what TestApplyFieldTags_JSONDash 테스트
package scanner

import "testing"

func TestApplyFieldTags_JSONDash(t *testing.T) {
	f := &Field{Name: "Secret"}
	excluded := applyFieldTags(f, `json:"-"`)
	if !excluded {
		t.Fatal("should be excluded")
	}
}
