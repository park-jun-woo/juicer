//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestNormalizeRole 테스트
package spring

import "testing"

func TestNormalizeRole(t *testing.T) {
	if normalizeRole("ROLE_ADMIN") != "ADMIN" || normalizeRole("USER") != "USER" {
		t.Fatal("normalize")
	}
}
