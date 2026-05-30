//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractClassRoute 테스트
package dotnet

import "testing"

func TestExtractClassRoute(t *testing.T) {
	root, src := parseCS(t, `[Route("api/[controller]")] class UsersController {}`)
	cls := findAllByType(root, "class_declaration")[0]
	got := extractClassRoute(cls, src, "Users")
	if got == "" {
		t.Fatal("expected non-empty route")
	}
}
