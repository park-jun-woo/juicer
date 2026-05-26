//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestBuildRegisteredSet_AnyMethod — Any 메서드가 5개 HTTP 메서드 모두 등록되는지 테스트
package scanner

import "testing"

func TestBuildRegisteredSet_AnyMethod(t *testing.T) {
	sr := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "ANY", Path: "/api/health"},
		},
	}
	set := buildRegisteredSet(sr)

	expected := []string{"get", "post", "put", "patch", "delete"}
	for _, m := range expected {
		key := m + "\t/api/health"
		if !set[key] {
			t.Errorf("expected %s to be registered for ANY", m)
		}
	}

	// head, options 등은 등록되지 않아야 함
	if set["head\t/api/health"] {
		t.Error("head should not be registered for ANY")
	}
}
