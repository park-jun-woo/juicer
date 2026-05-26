//ff:func feature=ddl type=test control=sequence
//ff:what TestExtractParenBody_UnclosedCov 테스트
package ddl

import "testing"

func TestExtractParenBody_UnclosedCov(t *testing.T) {
	got := extractParenBody("(id INT, name TEXT")
	if got != "id INT, name TEXT" {
		t.Fatalf("expected 'id INT, name TEXT', got %q", got)
	}
}
