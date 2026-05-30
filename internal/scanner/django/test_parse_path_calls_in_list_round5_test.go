//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParsePathCallsInList_Round5 테스트
package django

import "testing"

func TestParsePathCallsInList_Round5(t *testing.T) {
	src := []byte("urlpatterns = [path('a/', AView.as_view()), path('b/', BView.as_view())]")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	list := djFirst(t, root, "list")
	entries := parsePathCallsInList(list, src)
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
}
