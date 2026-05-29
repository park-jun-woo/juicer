//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestIsTestDir 테스트
package scanner

import "testing"

func TestIsTestDir(t *testing.T) {
	yes := []string{"test", "tests", "__tests__", "spec", "specs", "__mocks__", "fixtures", "testdata", "e2e", "Tests", "Spec"}
	for _, name := range yes {
		if !IsTestDir(name) {
			t.Errorf("IsTestDir(%q) = false, want true", name)
		}
	}
	no := []string{"src", "controllers", "routes", "internal", "api", "lib"}
	for _, name := range no {
		if IsTestDir(name) {
			t.Errorf("IsTestDir(%q) = true, want false", name)
		}
	}
}
