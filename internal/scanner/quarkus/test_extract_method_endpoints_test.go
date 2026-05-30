//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractMethodEndpoints 테스트
package quarkus

import "testing"

func TestExtractMethodEndpoints(t *testing.T) {
	fi := qFileInfo(t, `class R {
		@GET public String a() { return ""; }
		@POST public String b() { return ""; }
		public void helper() {}
	}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	eps := extractMethodEndpoints(cls, fi)
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}
