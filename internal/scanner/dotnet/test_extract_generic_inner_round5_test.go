//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractGenericInner_Round5 테스트
package dotnet

import "testing"

func TestExtractGenericInner_Round5(t *testing.T) {
	if got := extractGenericInner("List<Foo>"); got != "Foo" {
		t.Errorf("got %q", got)
	}
	if got := extractGenericInner("NoGeneric"); got != "NoGeneric" {
		t.Errorf("no-open: got %q", got)
	}
	if got := extractGenericInner("Broken<"); got != "Broken<" {
		t.Errorf("no-close: got %q", got)
	}
	if got := extractGenericInner(">Bad<"); got != ">Bad<" {
		t.Errorf("end<=start: got %q", got)
	}
}
