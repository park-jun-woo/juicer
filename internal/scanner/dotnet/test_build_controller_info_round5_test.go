//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestBuildControllerInfo_Round5 테스트
package dotnet

import "testing"

func TestBuildControllerInfo_Round5(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	cls := findAllByType(fi.root, "class_declaration")[0]
	ci := buildControllerInfo(cls, fi)
	if ci.className != "UsersController" {
		t.Fatalf("className: %q", ci.className)
	}
	if ci.prefix != "api/users" {
		t.Fatalf("prefix: %q", ci.prefix)
	}
	if len(ci.endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(ci.endpoints))
	}
}
