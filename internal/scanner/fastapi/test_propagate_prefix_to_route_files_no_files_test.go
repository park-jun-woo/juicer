//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagatePrefixToRouteFiles_NoFiles 테스트
package fastapi

import "testing"

func TestPropagatePrefixToRouteFiles_NoFiles(t *testing.T) {

	propagatePrefixToRouteFiles("/", nil)
}
