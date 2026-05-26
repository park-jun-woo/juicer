//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyClassDecorators_NoClassDecorators 테스트
package nestjs

import "testing"

func TestApplyClassDecorators_NoClassDecorators(t *testing.T) {
	ci := controllerInfo{
		endpoints: []endpointInfo{
			{handler: "findAll", middleware: []string{"LocalGuard"}},
		},
	}
	applyClassDecorators(&ci)
	if len(ci.endpoints[0].middleware) != 1 || ci.endpoints[0].middleware[0] != "LocalGuard" {
		t.Fatalf("expected [LocalGuard], got %v", ci.endpoints[0].middleware)
	}
}
