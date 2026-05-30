//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestBuildControllerInfoAndExtract 테스트
package dotnet

import "testing"

func TestBuildControllerInfoAndExtract(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	controllers := extractControllers(fi)
	if len(controllers) != 1 {
		t.Fatalf("expected 1, got %d", len(controllers))
	}
	if controllers[0].className != "UsersController" {
		t.Fatalf("name: %q", controllers[0].className)
	}
	if len(controllers[0].endpoints) != 2 {
		t.Fatalf("endpoints: %+v", controllers[0].endpoints)
	}
	got := collectControllers([]*fileInfo{fi})
	if len(got) != 1 {
		t.Fatalf("collect: %d", len(got))
	}
}
