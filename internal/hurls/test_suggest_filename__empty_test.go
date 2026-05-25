//ff:func feature=hurl type=parse control=sequence
//ff:what TestSuggestFilename_Empty 테스트
package hurls

import "testing"

func TestSuggestFilename_Empty(t *testing.T) {
	got := suggestFilename("/api/v1")
	if got != "test.hurl" {
		t.Fatalf("got %q", got)
	}
}
