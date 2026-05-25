//ff:func feature=hurl type=parse control=sequence
//ff:what TestSuggestFilename_Health 테스트
package hurls

import "testing"

func TestSuggestFilename_Health(t *testing.T) {
	got := suggestFilename("/api/health")
	if got != "health.hurl" {
		t.Fatalf("got %q", got)
	}
}
