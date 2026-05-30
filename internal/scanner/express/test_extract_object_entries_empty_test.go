//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractObjectEntries_Empty 테스트
package express

import "testing"

func TestExtractObjectEntries_Empty(t *testing.T) {
	fi := mustParse(t, []byte(`const m = [];`))
	entries := extractObjectEntries(firstArray(t, fi), fi.Src)
	if len(entries) != 0 {
		t.Fatalf("expected 0, got %d", len(entries))
	}
}
