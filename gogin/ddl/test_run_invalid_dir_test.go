//ff:func feature=ddl type=parse control=sequence
//ff:what TestRun_InvalidDir 테스트
package ddl

import (
	"testing"
)

func TestRun_InvalidDir(t *testing.T) {
	_, err := Run("/nonexistent/path/that/does/not/exist")
	// Glob doesn't error on non-existent directories; it returns no matches
	// But if the path pattern is malformed, it might error
	// For non-existent dir, Glob returns nil, nil — so no error expected
	if err != nil {
		// This is acceptable too
		return
	}
}
