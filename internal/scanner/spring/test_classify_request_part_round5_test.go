//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyRequestPart_Round5 테스트
package spring

import "testing"

func TestClassifyRequestPart_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@RequestPart("file") MultipartFile file) {} }`)
	var ep endpointInfo
	classifyRequestPart(param, src, &ep, "file")
	if len(ep.files) != 1 {
		t.Fatalf("files: %+v", ep.files)
	}
}
