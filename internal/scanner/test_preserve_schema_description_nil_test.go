//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveSchemaDescription_Nil 테스트
package scanner

import "testing"

func TestPreserveSchemaDescription_Nil(t *testing.T) {
	preserveSchemaDescription(nil, nil)
}
