//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestBuildResponse_RespTypeMapsEmpty 테스트
package fastapi

import "testing"

func TestBuildResponse_RespTypeMapsEmpty(t *testing.T) {

	resp := buildResponse(routeInfo{method: "GET"}, "None")
	_ = resp
}
