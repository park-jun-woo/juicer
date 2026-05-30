//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractControllersAndCollect 테스트
package spring

import "testing"

func TestExtractControllersAndCollect(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	controllers := extractControllers(fi)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	got := collectControllers([]*fileInfo{fi})
	if len(got) != 1 {
		t.Fatalf("collect: %d", len(got))
	}
}
