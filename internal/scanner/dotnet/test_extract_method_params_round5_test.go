//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractMethodParams_Round5 테스트
package dotnet

import "testing"

func TestExtractMethodParams_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M([FromQuery] string a, [FromBody] int b) {} }`)
	m := firstOfType(t, root, "method_declaration")
	var ep endpointInfo
	extractMethodParams(m, src, &ep)
	if len(ep.query) != 1 || ep.query[0].Name != "a" {
		t.Fatalf("query: %+v", ep.query)
	}
	if ep.bodyType == "" {
		t.Fatalf("expected body type set, got %+v", ep)
	}
}
