//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractUsingNamespace_Round5 테스트
package dotnet

import "testing"

func TestExtractUsingNamespace_Round5(t *testing.T) {
	root, src := parseCS(t, "using System.Text;\nclass C {}")
	using := firstOfType(t, root, "using_directive")
	if got := extractUsingNamespace(using, src); got != "System.Text" {
		t.Fatalf("got %q", got)
	}
}
