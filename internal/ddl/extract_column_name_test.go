//ff:func feature=ddl type=parse control=sequence
//ff:what TestExtractColumnName_Basic 테스트
package ddl

import "testing"

func TestExtractColumnName_Basic(t *testing.T) {
	got := extractColumnName("id INT PRIMARY KEY")
	if got != "id" {
		t.Fatalf("got %q", got)
	}
}
