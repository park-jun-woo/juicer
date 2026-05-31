//ff:func feature=scan type=test topic=dotnet control=sequence
//ff:what selectBodyResponse 2xx 우선 선택 및 첫 매칭 폴백 테스트
package dotnet

import "testing"

func TestSelectBodyResponse(t *testing.T) {
	// has both error and 2xx -> picks 2xx
	src := []byte(`class C { void M() {
		if (bad) return NotFound();
		return Ok(new Dto());
	} }`)
	root, _ := parseCSharp(src)
	body := findAllByType(root, "block")[0]
	res := selectBodyResponse(body, src)
	if !res.found || res.status != "200" {
		t.Errorf("should prefer 2xx: %+v", res)
	}

	// only non-2xx -> first found
	src2 := []byte(`class C { void M() { return NotFound(); } }`)
	root2, _ := parseCSharp(src2)
	body2 := findAllByType(root2, "block")[0]
	res2 := selectBodyResponse(body2, src2)
	if !res2.found || res2.status != "404" {
		t.Errorf("fallback first: %+v", res2)
	}
}
