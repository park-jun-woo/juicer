//ff:func feature=sql type=parse control=sequence
//ff:what TestRenderYAML 테스트
package sqls

import (
	"testing"
)

func TestRenderYAML(t *testing.T) {
	result := &SkeletonResult{
		Methods: []MethodSkeleton{
			{Repo: "UserRepo", Method: "GetAll", CRUD: "SELECT", Tables: []string{"users"}},
		},
	}
	out, err := RenderYAML(result)
	if err != nil {
		t.Fatalf("RenderYAML() error: %v", err)
	}
	if len(out) == 0 {
		t.Error("expected non-empty output")
	}
}
