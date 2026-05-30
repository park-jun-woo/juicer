//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyParam_RequestPart 테스트
package spring

import "testing"

func TestClassifyParam_RequestPart(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestPart("file") MultipartFile file) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.files) != 1 || ep.files[0].Type != "string:binary" {
		t.Fatalf("got %+v", ep.files)
	}
}
