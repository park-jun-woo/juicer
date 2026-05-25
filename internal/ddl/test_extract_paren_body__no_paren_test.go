//ff:func feature=ddl type=parse control=sequence
//ff:what TestExtractParenBody_NoParen 테스트
package ddl

import "testing"

func TestExtractParenBody_NoParen(t *testing.T) {
	got := extractParenBody("CREATE TABLE users")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}
