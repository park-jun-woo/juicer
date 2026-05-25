//ff:func feature=hurl type=parse control=sequence
//ff:what TestSuggestFilename_Basic 테스트
package hurls

import "testing"

func TestSuggestFilename_Basic(t *testing.T) {
	got := suggestFilename("/api/v1/admin/buildings")
	if got != "buildings.hurl" {
		t.Fatalf("got %q", got)
	}
}
