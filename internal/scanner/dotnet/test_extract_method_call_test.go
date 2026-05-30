//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestExtractMethodCall 테스트
package dotnet

import "testing"

func TestExtractMethodCall(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { app.MapGet("/x", h); } }`)
	invs := findAllByType(root, "invocation_expression")
	for _, inv := range invs {
		recv, method := extractMethodCall(inv, src)
		if recv == "app" && method == "MapGet" {
			return
		}
	}
	t.Fatal("did not find app.MapGet")
}
