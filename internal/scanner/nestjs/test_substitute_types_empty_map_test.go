//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestSubstituteTypes_EmptyMap 빈 타입맵으로 치환 시 원본 유지 테스트
package nestjs

import "testing"

func TestSubstituteTypes_EmptyMap(t *testing.T) {
	eps := []endpointInfo{
		{returnType: "D", bodyType: "CreateDto"},
	}
	substituteTypes(eps, map[string]string{})
	if eps[0].returnType != "D" {
		t.Errorf("expected no change, got returnType=%q", eps[0].returnType)
	}
	if eps[0].bodyType != "CreateDto" {
		t.Errorf("expected no change, got bodyType=%q", eps[0].bodyType)
	}
}
