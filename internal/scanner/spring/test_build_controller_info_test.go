//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestBuildControllerInfo 테스트
package spring

import "testing"

func TestBuildControllerInfo(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	cls := findAllByType(fi.root, "class_declaration")[0]
	ci := buildControllerInfo(cls, fi)
	if ci.className != "UserController" || ci.prefix != "/users" {
		t.Fatalf("meta: %+v", ci)
	}
	if len(ci.endpoints) != 2 {
		t.Fatalf("endpoints: %+v", ci.endpoints)
	}
}
