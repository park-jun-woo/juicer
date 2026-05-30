//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestBuildAllEndpoints_Round5 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildAllEndpoints_Round5(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	controllers := collectControllers([]*fileInfo{fi})
	if len(controllers) == 0 {
		t.Fatal("no controllers collected")
	}
	eps, _ := buildAllEndpoints(controllers, "/abs")
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d: %+v", len(eps), eps)
	}
	var _ scanner.Endpoint = eps[0]

	ceps, _ := buildControllerEndpoints(controllers[0], "/abs", 0)
	if len(ceps) != 2 {
		t.Fatalf("buildControllerEndpoints: expected 2, got %d", len(ceps))
	}
}
