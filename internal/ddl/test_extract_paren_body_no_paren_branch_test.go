//ff:func feature=ddl type=test control=sequence
//ff:what TestExtractParenBody_NoParenBranch 테스트
package ddl

import "testing"

func TestExtractParenBody_NoParenBranch(t *testing.T) {
	got := extractParenBody("CREATE TABLE users")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
