//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_FallbackCov 테스트
package scanner

import "testing"

func TestToYAMLNode_FallbackCov(t *testing.T) {
	toYAMLNode(3.14)
}
