//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveDescription_Nil 테스트
package scanner

import "testing"

func TestPreserveDescription_Nil(t *testing.T) {
	preserveDescription(nil, nil)
}

