//ff:func feature=hurl type=parse control=sequence
//ff:what TestSuggestFilename_WithParam 테스트
package hurls

import "testing"

func TestSuggestFilename_WithParam(t *testing.T) {
	got := suggestFilename("/api/v1/admin/buildings/:id")
	if got != "buildings_id.hurl" {
		t.Fatalf("got %q", got)
	}
}
