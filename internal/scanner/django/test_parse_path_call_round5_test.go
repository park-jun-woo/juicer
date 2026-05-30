//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParsePathCall_Round5 테스트
package django

import "testing"

func TestParsePathCall_Round5(t *testing.T) {
	src := []byte("urlpatterns = [path('users/', UserView.as_view())]")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := djFirst(t, root, "call")
	entry := parsePathCall(call, src)
	if entry == nil {
		t.Fatal("expected url entry")
	}
	if entry.pattern != "users/" {
		t.Errorf("pattern: %q", entry.pattern)
	}
	if entry.viewName != "UserView" {
		t.Errorf("viewName: %q", entry.viewName)
	}
}
