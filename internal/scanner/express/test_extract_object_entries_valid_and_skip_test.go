//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractObjectEntries_ValidAndSkip 테스트
package express

import "testing"

func TestExtractObjectEntries_ValidAndSkip(t *testing.T) {

	fi := mustParse(t, []byte(`const m = [ { path: '/a', route: userRoute }, { path: '/b' } ];`))
	entries := extractObjectEntries(firstArray(t, fi), fi.Src)
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d (%v)", len(entries), entries)
	}
	if entries[0].path != "/a" || entries[0].routeVar != "userRoute" {
		t.Fatalf("unexpected entry %+v", entries[0])
	}
}
