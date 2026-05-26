//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveAllDTOs_Empty 테스트
package nestjs

import "testing"

func TestResolveAllDTOs_Empty(t *testing.T) {
	resolveAllDTOs(nil, nil)
}
