//ff:func feature=ddl type=parse control=sequence
//ff:what TestCleanLine_WithComment 테스트
package ddl

import "testing"

func TestCleanLine_WithComment(t *testing.T) {
	got := cleanLine("  id INT NOT NULL -- primary key  ")
	if got != "id INT NOT NULL" {
		t.Fatalf("got %q", got)
	}
}
