//ff:func feature=sql type=parse control=sequence
//ff:what TestRenderJSON 테스트
package sqls

import (
	"testing"
)

func TestRenderJSON(t *testing.T) {
	result := &SkeletonResult{
		Methods: []MethodSkeleton{
			{Repo: "UserRepo", Method: "GetAll", CRUD: "SELECT", Tables: []string{"users"}},
		},
	}
	out, err := RenderJSON(result)
	if err != nil {
		t.Fatalf("RenderJSON() error: %v", err)
	}
	if len(out) == 0 {
		t.Error("expected non-empty output")
	}
}
