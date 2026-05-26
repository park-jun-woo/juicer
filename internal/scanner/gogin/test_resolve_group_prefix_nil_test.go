//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefix_NilPkgs 테스트
package gogin

import (
	"testing"
)

func TestResolveGroupPrefix_NilPkgs(t *testing.T) {
	// nil inputs should not panic
	resolveGroupPrefix(nil, ".", nil, nil, nil)
}
