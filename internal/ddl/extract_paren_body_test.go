//ff:func feature=ddl type=parse control=sequence
//ff:what TestExtractParenBody_Basic 테스트
package ddl

import "testing"

func TestExtractParenBody_Basic(t *testing.T) {
	got := extractParenBody("CREATE TABLE users (id INT, name TEXT)")
	if got != "id INT, name TEXT" {
		t.Fatalf("got %q", got)
	}
}
