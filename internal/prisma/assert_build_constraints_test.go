//ff:func feature=prisma type=test control=iteration dimension=1
//ff:what 주어진 모델/타깃으로 buildConstraints 결과가 기대값과 같은지 검증하는 테스트 헬퍼
package prisma

import (
	"reflect"
	"testing"
)

// assertBuildConstraints resolves target within models and checks that
// buildConstraints returns want.
func assertBuildConstraints(t *testing.T, models []model, target string, want []string) {
	t.Helper()
	s := buildSchema(models)
	var tm model
	for _, m := range models {
		if m.name == target {
			tm = m
		}
	}
	got := buildConstraints(tm, s)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildConstraints(%s) =\n  %#v\nwant\n  %#v", target, got, want)
	}
}
