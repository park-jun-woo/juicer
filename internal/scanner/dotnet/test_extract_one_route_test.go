//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractOneRoute 테스트
package dotnet

import "testing"

func TestExtractOneRoute(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep, ok := extractOneRoute(m, fi)
	if !ok || ep.method != "GET" || ep.handler != "Get" {
		t.Fatalf("got %+v ok=%v", ep, ok)
	}
}
