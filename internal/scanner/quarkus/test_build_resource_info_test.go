//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildResourceInfo 테스트
package quarkus

import "testing"

func TestBuildResourceInfo(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	cls := findAllByType(fi.root, "class_declaration")[0]
	ri := buildResourceInfo(cls, fi)
	if ri.className != "UserResource" || ri.prefix != "/users" {
		t.Fatalf("meta: %+v", ri)
	}
	if len(ri.roles) != 1 || len(ri.endpoints) != 2 {
		t.Fatalf("roles/endpoints: %+v", ri)
	}
}
