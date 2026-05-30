//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractMapGroupsFromFile_Round5 테스트
package dotnet

import "testing"

func TestExtractMapGroupsFromFile_Round5(t *testing.T) {
	fi := csFileInfo(t, `
var app = builder.Build();
var api = app.MapGroup("/api");
`)
	groups := map[string]string{}
	extractMapGroupsFromFile(fi, groups)
	if groups["api"] != "/api" {
		t.Fatalf("expected api=/api, got %v", groups)
	}
}
