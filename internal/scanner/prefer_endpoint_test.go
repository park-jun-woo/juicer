//ff:func feature=scan type=test control=sequence topic=scanner
//ff:what preferEndpoint 테스트 (round5)
package scanner

import "testing"

func TestPreferEndpoint_Round5(t *testing.T) {
	rich := Endpoint{
		Path:      "/a",
		Responses: []Response{{TypeName: "User"}},
	}
	poor := Endpoint{Path: "/a"}

	// richer candidate wins
	if !preferEndpoint(rich, poor) {
		t.Error("richer candidate should be preferred")
	}
	if preferEndpoint(poor, rich) {
		t.Error("poorer candidate should not be preferred")
	}

	// tie on richness => smaller File wins
	a := Endpoint{File: "a.go", Line: 10}
	b := Endpoint{File: "b.go", Line: 1}
	if !preferEndpoint(a, b) {
		t.Error("smaller File should win on richness tie")
	}
	if preferEndpoint(b, a) {
		t.Error("larger File should not win")
	}

	// tie on richness and File => smaller Line wins
	c := Endpoint{File: "x.go", Line: 5}
	dd := Endpoint{File: "x.go", Line: 9}
	if !preferEndpoint(c, dd) {
		t.Error("smaller Line should win")
	}
	if preferEndpoint(dd, c) {
		t.Error("larger Line should not win")
	}
}
