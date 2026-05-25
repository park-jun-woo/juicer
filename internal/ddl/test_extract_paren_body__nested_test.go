//ff:func feature=ddl type=parse control=sequence
//ff:what TestExtractParenBody_Nested 테스트
package ddl

import "testing"

func TestExtractParenBody_Nested(t *testing.T) {
	got := extractParenBody("CREATE TABLE t (id INT CHECK(id > 0), name TEXT)")
	if got != "id INT CHECK(id > 0), name TEXT" {
		t.Fatalf("got %q", got)
	}
}
