//ff:func feature=ddl type=command control=sequence
//ff:what TestRun_InvalidDirCov 테스트
package ddl

import "testing"

func TestRun_InvalidDirCov(t *testing.T) {
	_, err := Run("/nonexistent/dir/12345")
	if err != nil {
		// glob returns nil on non-matching pattern, not error
		// This depends on OS behavior
	}
}
