//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractMinimalAPIsFromFile_Round5 테스트
package dotnet

import "testing"

func TestExtractMinimalAPIsFromFile_Round5(t *testing.T) {
	fi := csFileInfo(t, `
var app = builder.Build();
app.MapGet("/health", () => Results.Ok());
`)
	eps := extractMinimalAPIsFromFile(fi, map[string]string{})
	if len(eps) != 1 || eps[0].Path != "/health" {
		t.Fatalf("got %+v", eps)
	}
}
